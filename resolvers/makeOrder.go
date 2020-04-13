package resolvers

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/3dw1nM0535/deli/db"
	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
)

func dishExists(id string) bool {
	dish := models.Dish{}
	uid := "00000000-0000-0000-0000-000000000000"
	parseID := utils.ParseUUID(id)
	if parseID.String() == uid {
		return false
	}
	db, _ := db.Factory()
	db.DB.Where("id = ?", parseID).First(&dish)
	if dish.ID.String() == uid {
		return false
	}
	return true
}

func mapItemsToNotes(items []*models1.DishNote) ([]*models.DishOrder, error) {
	var dishes = []*models.DishOrder{}
	var err error

	for i := range items {
		// validate input
		if items[i].DishID == "" {
			err = errors.New("dish id cannot be empty")
			return dishes, err
		}
		if dishExists(items[i].DishID) == false {
			err = errors.New("dish does not exist")
			return dishes, err
		}
		if items[i].Title == "" {
			err = errors.New("dish title cannot be empty")
			return dishes, err
		}
		if items[i].Description == "" {
			err = errors.New("dish description cannot be empty")
			return dishes, err
		}
		dI := &models.DishOrder{
			DishID:      utils.ParseUUID(items[i].DishID),
			Title:       items[i].Title,
			Description: items[i].Description,
			AddOns:      items[i].AddOns,
			Price:       items[i].Price,
			Count:       items[i].Count,
			Subtotal:    items[i].Subtotal,
			OrderID:     utils.ParseUUID(models.ID),
		}
		dishes = append(dishes, dI)
	}
	return dishes, err
}

func (r *mutationResolver) MakeOrder(ctx context.Context, input models1.OrderInput) (*models.Order, error) {
	restaurant := &models.Restaurant{}
	r.ORM.DB.Where("id = ?", input.RestaurantID).First(&restaurant)
	if restaurant.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("no restaurant with id %s", input.RestaurantID)
		return &models.Order{}, err
	}
	// Validate user phone number
	if input.PhoneNumber == "" {
		return &models.Order{}, errors.New("You must provide your registered Mpesa phone number to complete your order")
	}

	// Validate and confirm payment
	totalSum := strconv.Itoa(int(input.TotalSum))

	res := utils.MakePayment(totalSum, input.PhoneNumber)
	checkoutRes, _, statusCode := utils.ConfirmPayment(res)
	// validate pin
	if statusCode == 200 && checkoutRes.ResultDesc == utils.MpesaWrongPinMessage {
		return &models.Order{}, errors.New("Wrong Mpesa Pin")
	}
	// validate insufficient account balance for payment
	if statusCode == 200 && checkoutRes.ResultDesc == utils.MpesaInsufficientMessage {
		return &models.Order{}, errors.New(checkoutRes.ResultDesc)
	}
	// validate timeout request
	if statusCode == 200 && checkoutRes.ResultDesc == utils.MpesaTimeoutRequestMessage {
		return &models.Order{}, errors.New("You cancelled the payment")
	}
	// validate successful payment
	if statusCode == 200 && checkoutRes.ResultDesc == utils.MpesaSuccessfulRequestMessage {
		order := &models.Order{
			RestaurantNotes: input.RestaurantNotes,
			RestaurantID:    utils.ParseUUID(input.RestaurantID),
			TotalSum:        input.TotalSum,
			OrderStatus:     "In-Kitchen",
		}
		r.ORM.DB.Create(&order)
		payment := &models.Payment{
			MerchantRequestID: checkoutRes.MerchantRequestID,
			CheckoutRequestID: checkoutRes.CheckoutRequestID,
			MpesaDescription:  checkoutRes.ResultDesc,
			OrderID:           utils.ParseUUID(models.ID),
			OrderPaidFor:      order,
			RestaurantID:      utils.ParseUUID(input.RestaurantID),
			PaidAmount:        input.TotalSum,
			PhoneNumber:       input.PhoneNumber,
			Confirmed:         true,
		}
		r.ORM.DB.Save(&payment)

		dishes, err := mapItemsToNotes(input.OrderNotes)
		if err != nil {
			return &models.Order{}, err
		}

		for i := range dishes {
			r.ORM.DB.Save(dishes[i])
		}

		for _, c := range orderCreatedChannel[input.RestaurantID] {
			c <- order
		}

		return order, nil
	}
	return &models.Order{}, errors.New(checkoutRes.ResultDesc)
}
