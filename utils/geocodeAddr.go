package utils

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/deli/db/models"
	"googlemaps.github.io/maps"
)

// GeoCodeAddr : return geocoded restaurant address
func GeoCodeAddr(ctx context.Context, address *models.Address, apiKey string) ([]maps.GeocodingResult, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	addr := &maps.GeocodingRequest{
		Address:    fmt.Sprintf("%s %s", address.StreetName, address.PostalTown),
		Components: map[maps.Component]string{"locality": address.PostalTown},
		Region:     "KE",
	}
	res, err := c.Geocode(ctx, addr)
	if err != nil {
		return nil, err
	}
	return res, nil
}
