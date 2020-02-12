package main

import (
	"github.com/3dw1nM0535/deli/db"
	"github.com/3dw1nM0535/deli/db/models"
)

func migrate() error {
	dbm, err := db.Factory()

	defer dbm.Close()

	err = dbm.AutoMigrate(&models.Deli{}).Error
	if err != nil {
		return err
	}
	return nil
}

func main() {
	migrate()
}
