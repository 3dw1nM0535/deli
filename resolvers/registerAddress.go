package resolvers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	models1 "github.com/3dw1nM0535/deli/db/models"
	"github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
)

func (r *mutationResolver) RegisterAddress(ctx context.Context, input models.AddressInput) (*models1.Address, error) {
	addr := &models1.Address{}
	if input.RestaurantID == "" {
		return &models1.Address{}, errors.New("restaurant id cannot be empty")
	}
	if input.PostalCode == "" {
		return &models1.Address{}, errors.New("postal code cannot be empty")
	}
	if input.City == "" {
		return &models1.Address{}, errors.New("provide your business primary city/town location")
	}
	r.ORM.DB.Raw("SELECT * FROM postals WHERE postal_code = ?", strings.Title(string(input.PostalCode))).Scan(&addr)
	if addr.PostalTown != "" {
		var restaurant = &models1.Restaurant{}
		r.ORM.DB.First(&restaurant, "id = ?", input.RestaurantID)
		if restaurant.ID.String() == "00000000-0000-0000-0000-000000000000" {
			err := fmt.Errorf("no restaurant with id '%s' to attach address to", input.RestaurantID)
			return &models1.Address{}, err
		}
		var restaurants = []*models1.Restaurant{restaurant}
		var address = &models1.Address{
			PostalCode:  input.PostalCode,
			PostalTown:  addr.PostalTown,
			City:        input.City,
			StreetName:  input.StreetName,
			Restaurants: restaurants,
		}
		geoCode, err := utils.GeoCodeAddr(context.Background(), address, restaurant.RestaurantName, geocodingKey)
		if err != nil {
			return &models1.Address{}, err
		}
		if len(geoCode) > 0 {
			for i := range geoCode {
				geoAddr := &models1.Address{
					PostalCode:    input.PostalCode,
					PostalTown:    addr.PostalTown,
					City:          input.City,
					StreetName:    input.StreetName,
					Lon:           geoCode[i].Geometry.Location.Lng,
					Lat:           geoCode[i].Geometry.Location.Lat,
					AddressString: geoCode[i].FormattedAddress,
					Restaurants:   restaurants,
				}
				if geoAddr.StreetName == "" {
					geoAddr.StreetName = geoCode[i].PlusCode.CompoundCode
				}
				r.ORM.DB.Save(&geoAddr)
				return geoAddr, nil
			}
		}
		if address.StreetName == "" {
			address.StreetName = geoCode[0].PlusCode.CompoundCode
		}
		address.Lon = geoCode[0].Geometry.Location.Lng
		address.Lat = geoCode[0].Geometry.Location.Lat
		address.AddressString = geoCode[0].FormattedAddress
		r.ORM.DB.Save(&address)
		return address, nil
	}
	err := fmt.Errorf("postal code '%s' doesn't exist", input.PostalCode)
	return &models1.Address{}, err
}
