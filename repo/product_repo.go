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

func GetProductsByIds(ids []int) ([]ProductDb, error) {
	var products []ProductDb
	err := getDb().Where("id IN ?", ids).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
