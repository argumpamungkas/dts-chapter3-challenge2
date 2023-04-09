package repo

import (
	"DTS/Chapter-3/chapter3-challenge2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "gume98"
	port     = "5432"
	dbname   = "api-product"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Println("Invalid connect database")
		return
	}

	log.Println("Success connect to database")

	db.Debug().AutoMigrate(models.Role{}, models.User{}, models.Product{})
	// db.Debug().Migrator().DropTable(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
