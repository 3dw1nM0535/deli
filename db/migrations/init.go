package main

import (
	"github.com/3dw1nM0535/Byte/db"
	"github.com/3dw1nM0535/Byte/db/models"
)

func migrate() error {
	orm, err := db.Factory()

	defer orm.DB.Close()

	err = orm.DB.AutoMigrate(
		&models.Farm{},
		&models.Season{},
	).Error
	if err != nil {
		return err
	}

	// Add geolocation column of type geography
	//orm.DB.Exec("ALTER TABLE addresses ADD COLUMN geolocation geography(point);")

	return nil
}

func main() {
	migrate()
}
