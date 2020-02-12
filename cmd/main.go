package main

import (
	"fmt"

	"github.com/3dw1nM0535/deli/db"
	"github.com/3dw1nM0535/deli/server"
)

func main() {
	orm, err := db.Factory()
	if err != nil {
		fmt.Errorf("Error connecting to database: %v", err.Error())
	}
	server.Run(orm)
}
