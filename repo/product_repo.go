package repo

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDb() *gorm.DB {
	DATABASE_URL := "postgres://kolaczyn@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})
	if err != nil {
		// it doesn't make sense to continue if we can't connect to the database
		panic(err)
	}

	db.AutoMigrate(&ProductDb{})

	return db
}

func SeedDb() {
	db := getDb()

	db.Create(&ProductDb{
		Id:    1,
		Name:  "Product 1",
		Price: 1.99,
	})
	db.Create(&ProductDb{
		Id:    2,
		Name:  "Product 2",
		Price: 2.99,
	})
	db.Create(&ProductDb{
		Id:    3,
		Name:  "Product 3",
		Price: 3.99,
	})

}

func GetAllProducts() []ProductDb {
	db := getDb()

	var products []ProductDb
	db.Find(&products)

	for _, p := range products {
		fmt.Println(p.Id)
	}

	return products
}
