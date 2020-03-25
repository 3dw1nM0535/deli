package db

import (
	"fmt"

	"github.com/3dw1nM0535/deli/utils"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbHost, dbPort, dbUser, dbPassword, dbName string
var sslMode bool

func init() {
	godotenv.Load()
	dbHost = utils.MustGetEnv("DBHOST")
	dbPort = utils.MustGetEnv("DBPORT")
	dbUser = utils.MustGetEnv("DBUSER")
	dbPassword = utils.MustGetEnv("DBPASS")
	dbName = utils.MustGetEnv("DBNAME")
	sslMode = utils.MustGetEnvBool("SSLMODE_ENABLED")

}

// DB : database structure
type DB struct {
	*gorm.DB
}

// Factory : open a database connection
func Factory() (*DB, error) {
	dbm, err := gorm.Open("postgres",
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	if err != nil {
		fmt.Printf("Error connecting to database: " + err.Error())
	}

	return &DB{dbm}, nil
}
