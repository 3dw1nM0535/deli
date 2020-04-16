package resolvers

import (
	"context"

	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
)

func (r *orderResolver) ID(ctx context.Context, obj *models.Order) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *restaurantResolver) ID(ctx context.Context, obj *models.Restaurant) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *dishResolver) ID(ctx context.Context, obj *models.Dish) (string, error) {
	id := obj.ID.String()
	return id, nil
}

// func (r *fileResolver) ID(ctx context.Context, obj *models.File) (string, error) {
// 	id := obj.ID.String()
// 	return id, nil
// }

func (r *dishResolver) AddOns(ctx context.Context, obj *models.Dish) ([]string, error) {
	var addOns []string
	addOns = obj.AddOns
	return addOns, nil
}

func (r *menuResolver) ID(ctx context.Context, obj *models.Menu) (string, error) {
	id := obj.ID.String()
	return id, nil
}

// Dishes : find dishes belonging to a menu
func (r *menuResolver) Dishes(ctx context.Context, obj *models.Menu) ([]*models.Dish, error) {
	dishes := []*models.Dish{}
	menu := obj
	r.ORM.DB.First(&menu, "id = ?", obj.ID)
	r.ORM.DB.Model(&menu).Related(&dishes)
	return dishes, nil
}

// Addresses : find addresses belonging to a restaurant
func (r *restaurantResolver) Addresses(ctx context.Context, obj *models.Restaurant) ([]*models.Address, error) {
	addresses := []*models.Address{}
	restaurant := obj
	r.ORM.DB.First(&restaurant, "id = ?", obj.ID)
	r.ORM.DB.Model(&restaurant).Related(&addresses, "Address")
	return addresses, nil
}

// Restaurants : find restaurants belonging to an address
func (r *addressResolver) Restaurants(ctx context.Context, obj *models.Address) ([]*models.Restaurant, error) {
	restaurants := []*models.Restaurant{}
	address := obj
	r.ORM.DB.First(&address, "id = ?", obj.ID)
	r.ORM.DB.Model(&address).Related(&restaurants, "Restaurants")
	return restaurants, nil
}

// License : find license belonging to a restaurant
func (r *restaurantResolver) License(ctx context.Context, obj *models.Restaurant) (*models1.File, error) {
	license := &models.License{}
	restaurant := obj
	r.ORM.DB.First(&restaurant, "id = ?", obj.ID)
	r.ORM.DB.Model(&obj).Related(&license, "License")
	return &models1.File{
		ID:        license.ID.String(),
		Media:     license.Media,
		Content:   license.Content,
		Size:      int(license.Size),
		CreatedAt: &license.CreatedAt,
		UpdatedAt: &license.UpdatedAt,
	}, nil
}

// Menu : find menu belonging to a restaurant
func (r *restaurantResolver) Menu(ctx context.Context, obj *models.Restaurant) ([]*models.Menu, error) {
	menu := []*models.Menu{}
	restaurant := obj
	r.ORM.DB.First(&restaurant, "id = ?", obj.ID)
	r.ORM.DB.Model(&restaurant).Related(&menu)
	return menu, nil
}

func (r *dishOrderResolver) ID(ctx context.Context, obj *models.DishOrder) (string, error) {
	id := obj.ID.String()
	return id, nil
}
func (r *dishOrderResolver) AddOns(ctx context.Context, obj *models.DishOrder) ([]string, error) {
	addons := obj.AddOns
	return addons, nil
}

// Notes : find notes belonging to an order
func (r *orderResolver) Notes(ctx context.Context, obj *models.Order) ([]*models.DishOrder, error) {
	notes := []*models.DishOrder{}
	order := obj
	r.ORM.DB.First(&order, "id = ?", order.ID)
	r.ORM.DB.Model(&order).Related(&notes)
	return notes, nil
}

// Orders : find order belonging to a restaurant
func (r *restaurantResolver) Orders(ctx context.Context, obj *models.Restaurant) ([]*models.Order, error) {
	orders := []*models.Order{}
	restaurant := obj
	r.ORM.DB.First(&restaurant, "id = ?", obj.ID)
	r.ORM.DB.Model(&restaurant).Related(&orders)
	return orders, nil
}

// Payment : find payment belonging to an order
func (r *orderResolver) Payment(ctx context.Context, obj *models.Order) (*models.Payment, error) {
	payment := &models.Payment{}
	order := obj
	r.ORM.DB.First(&order, "id = ?", obj.ID)
	r.ORM.DB.Model(&order).Related(&payment)
	return payment, nil
}

func (r *paymentResolver) ID(ctx context.Context, obj *models.Payment) (string, error) {
	id := obj.ID.String()
	return id, nil
}

// Payments : find payments belonging to a restaurant
func (r *restaurantResolver) Payments(ctx context.Context, obj *models.Restaurant) ([]*models.Payment, error) {
	payments := []*models.Payment{}
	restaurant := obj
	r.ORM.DB.First(&restaurant, "id = ?", obj.ID)
	r.ORM.DB.Model(&restaurant).Related(&payments)
	return payments, nil
}

func (r *riderResolver) ID(ctx context.Context, obj *models.Rider) (string, error) {
	id := obj.ID.String()
	return id, nil
}

// IdentificationDocument : find id belonging to a rider
func (r *riderResolver) IdentificationDocument(ctx context.Context, obj *models.Rider) (*models1.File, error) {
	idd := &models.IDD{}
	rider := obj
	r.ORM.DB.First(&rider, "id = ?", obj.ID)
	r.ORM.DB.Model(&rider).Related(&idd)
	return &models1.File{
		ID:        idd.ID.String(),
		Media:     idd.Media,
		Content:   idd.Content,
		CreatedAt: &idd.CreatedAt,
		UpdatedAt: &idd.UpdatedAt,
	}, nil
}

// MedicalCertificate : find medical certificate belonging to a rider
func (r *riderResolver) MedicalCertificate(ctx context.Context, obj *models.Rider) (*models1.File, error) {
	rider := obj
	mdc := &models.MDC{}
	r.ORM.DB.First(&rider, "id = ?", obj.ID)
	r.ORM.DB.Model(&rider).Related(&mdc)
	return &models1.File{
		ID:        mdc.ID.String(),
		Media:     mdc.Media,
		Content:   mdc.Content,
		CreatedAt: &mdc.CreatedAt,
		UpdatedAt: &mdc.UpdatedAt,
	}, nil
}

// GoodConductCertificate : find good conduct certificate belonging to a rider
func (r *riderResolver) GoodConductCertificate(ctx context.Context, obj *models.Rider) (*models1.File, error) {
	rider := obj
	gcc := &models.GCC{}
	r.ORM.DB.First(&rider, "id = ?", obj.ID)
	r.ORM.DB.Model(&rider).Related(&gcc)
	return &models1.File{
		ID:        gcc.ID.String(),
		Media:     gcc.Media,
		Content:   gcc.Content,
		CreatedAt: &gcc.CreatedAt,
		UpdatedAt: &gcc.UpdatedAt,
	}, nil
}
