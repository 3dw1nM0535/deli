package resolvers

import (
	"context"
	"errors"

	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
)

func (r *queryResolver) FindNearByRestaurants(ctx context.Context, input models1.Cords) ([]*models.Restaurant, error) {
	var lon float64 = input.Lon
	var lat float64 = input.Lat
	var restaurants []*models.Restaurant

	// validate provided cords
	if !(lon >= -180 && lon <= 180) {
		return []*models.Restaurant{}, errors.New("longitude should be a value between -180 and 180")
	}
	if !(lat >= -90 && lat <= 90) {
		return []*models.Restaurant{}, errors.New("latitude should be a value between -90 and 90")
	}

	nearByQuery := `SELECT *, ST_Distance(geolocation, ST_SetSRID(ST_MakePoint(?, ?), 4326)) FROM restaurants INNER JOIN restaurant_addresses ON restaurants.id = restaurant_addresses.restaurant_id INNER JOIN addresses ON addresses.id = restaurant_addresses.address_id WHERE ST_DWithin(geolocation, ST_SetSRID(ST_MakePoint(?, ?), 4326), 1000) ORDER BY st_distance ASC;`
	var dict []*struct{ StDistance float64 }
	r.ORM.DB.Raw(nearByQuery, lon, lat, lon, lat).Scan(&restaurants).Scan(&dict)
	for i := range restaurants {
		restaurants[i].Distance = dict[i].StDistance
	}
	return restaurants, nil
}
