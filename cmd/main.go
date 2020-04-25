package main

import (
	"log"

	"github.com/3dw1nM0535/Byte/db"
	"github.com/3dw1nM0535/Byte/server"
)

func main() {
	orm, err := db.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm)
}
