package main

import (
	"github.com/3dw1nM0535/deli/db"
	"github.com/3dw1nM0535/deli/db/models"
)

func migrate() error {
	orm, err := db.Factory()

	defer orm.DB.Close()

	err = orm.DB.AutoMigrate(&models.Restaurant{}, &models.Address{}).Error
	if err != nil {
		return err
	}
	return nil
}

func main() {
	migrate()
}
