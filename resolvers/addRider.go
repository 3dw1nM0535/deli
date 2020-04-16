package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
)

func (r *mutationResolver) AddRider(ctx context.Context, input models1.RiderInput) (*models.Rider, error) {
	rider := &models.Rider{}

	// validate input
	if input.EmailAddress == "" {
		return &models.Rider{}, errors.New("email address cannot be empty")
	}
	if input.Firstname == "" {
		return &models.Rider{}, errors.New("rider first name cannot be empty")
	}
	if input.Lastname == "" {
		return &models.Rider{}, errors.New("rider last name cannot be empty")
	}
	if input.PhoneNumber == "" {
		return &models.Rider{}, errors.New("rider phone number cannot be empty")
	}
	if len(input.PhoneNumber) < 12 || len(input.PhoneNumber) > 12 {
		return &models.Rider{}, errors.New("invalid phone number format")
	}

	// check if email is unique
	r.ORM.DB.First(&rider, "email_address = ?", input.EmailAddress)
	if rider.ID.String() != "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("rider with email '%s' already exists", input.EmailAddress)
		return &models.Rider{}, err
	}

	// check if phone_number is unique
	r.ORM.DB.First(&rider, "phone_number = ?", input.PhoneNumber)
	if rider.ID.String() != "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("rider with phone number '%s' already exists", input.PhoneNumber)
		return &models.Rider{}, err
	}

	newRider := &models.Rider{
		FirstName:    input.Firstname,
		LastName:     input.Lastname,
		EmailAddress: input.EmailAddress,
		PhoneNumber:  input.PhoneNumber,
	}
	r.ORM.DB.Save(&newRider)
	return newRider, nil
}
