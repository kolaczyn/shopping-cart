package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDb() *gorm.DB {
	DATABASE_URL := "postgres://kolaczyn@localhost:5432/postgres"

	// TODO rename the package from db to repository and rename conn to db
	conn, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})
	if err != nil {
		// it doesn't make sense to continue if we can't connect to the database
		panic(err)
	}

	conn.AutoMigrate(&ProductDb{})

	return conn
}

func SeedDb() {
	conn := getDb()

	conn.Create(&ProductDb{
		Id:    1,
		Name:  "Product 1",
		Price: 1.99,
	})
	conn.Create(&ProductDb{
		Id:    2,
		Name:  "Product 2",
		Price: 2.99,
	})
	conn.Create(&ProductDb{
		Id:    3,
		Name:  "Product 3",
		Price: 3.99,
	})

}

func GetAllProducts() []ProductDb {
	conn := getDb()

	var products []ProductDb
	conn.Find(&products)

	for _, p := range products {
		fmt.Println(p.Id)
	}

	return products
}
