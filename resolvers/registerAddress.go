package resolvers

import (
	"context"
	"fmt"
	"log"
	"strings"

	models1 "github.com/3dw1nM0535/deli/db/models"
	"github.com/3dw1nM0535/deli/models"
)

func (r *mutationResolver) RegisterAddress(ctx context.Context, input models.AddressInput) (*models1.Address, error) {
	addr := &models1.Address{}
	r.ORM.DB.Raw("SELECT postal_town FROM postals WHERE postal_town = ?", strings.Title(string(input.PostalTown))).Scan(&addr)
	log.Println(addr.PostalTown)
	if addr.PostalTown != "" {
		var address = &models1.Address{
			PostalCode:   input.PostalCode,
			PostalTown:   input.PostalTown,
			StreetName:   input.StreetName,
			BuildingName: input.BuildingName,
			Lon:          input.Lon,
			Lat:          input.Lat,
		}
		r.ORM.DB.Create(&address)
		return address, nil
	}
	err := fmt.Errorf("postal town '%s' doesn't exist", input.PostalTown)
	return &models1.Address{}, err
}

// func (r *mutationResolver) RegisterAddress(ctx context.Context, input models.AddressInput) (*models1.Address, error) {

// }
