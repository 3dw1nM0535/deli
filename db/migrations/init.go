package main

import (
	"github.com/3dw1nM0535/deli/db"
	"github.com/3dw1nM0535/deli/db/models"
)

func migrate() error {
	orm, err := db.Factory()

	defer orm.DB.Close()

	err = orm.DB.AutoMigrate(
		&models.Restaurant{},
		&models.Address{},
		&models.License{},
		&models.Menu{},
		&models.Dish{},
		&models.DishOrder{},
		&models.Order{},
		&models.Payment{},
		&models.IDD{},
		&models.MDC{},
		&models.GCC{},
		&models.Rider{},
		&models.DisplayPicture{},
	).Error
	if err != nil {
		return err
	}

	// Add foreign keys
	orm.DB.Model(&models.RestaurantAddresses{}).AddForeignKey("restaurant_id", "restaurants(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.RestaurantAddresses{}).AddForeignKey("address_id", "addresses(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.License{}).AddForeignKey("restaurant_id", "restaurants(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.Menu{}).AddForeignKey("restaurant_id", "restaurants(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.Dish{}).AddForeignKey("menu_id", "menus(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.DishOrder{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.DishOrder{}).AddForeignKey("dish_id", "dishes(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.Order{}).AddForeignKey("restaurant_id", "restaurants(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.Payment{}).AddForeignKey("restaurant_id", "restaurants(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.Payment{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.Rider{}).AddUniqueIndex("idx_rider_email", "email_address")
	orm.DB.Model(&models.Rider{}).AddUniqueIndex("idx_rider_phone_no", "phone_number")
	orm.DB.Model(&models.IDD{}).AddForeignKey("rider_id", "riders(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.MDC{}).AddForeignKey("rider_id", "riders(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.GCC{}).AddForeignKey("rider_id", "riders(id)", "CASCADE", "CASCADE")
	orm.DB.Model(&models.DisplayPicture{}).AddForeignKey("rider_id", "riders(id)", "CASCADE", "CASCADE")

	// Add geolocation column of type geography
	orm.DB.Exec("ALTER TABLE addresses ADD COLUMN geolocation geography(point);")

	return nil
}

func main() {
	migrate()
}
