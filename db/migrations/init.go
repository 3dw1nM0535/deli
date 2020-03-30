package main

import (
	"github.com/3dw1nM0535/deli/db"
	"github.com/3dw1nM0535/deli/db/models"
)

func migrate() error {
	orm, err := db.Factory()

	defer orm.DB.Close()

	orm.DB.DropTableIfExists(&models.RestaurantAddresses{}, &models.Restaurant{}, &models.Address{})

	err = orm.DB.AutoMigrate(&models.Restaurant{}, &models.Address{}).Error
	if err != nil {
		return err
	}

	// Add foreign keys
	orm.DB.Model(&models.RestaurantAddresses{}).AddForeignKey("restaurant_id", "restaurants(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.RestaurantAddresses{}).AddForeignKey("address_id", "addresses(id)", "CASCADE", "CASCADE")

	return nil
}

func main() {
	migrate()
}
