package repo

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDb() *gorm.DB {
	dbUrl := os.Getenv("SHOPPING_PRODUCT_DB_URL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		// it doesn't make sense to continue if we can't connect to the database
		panic(err)
	}

	return db
}
