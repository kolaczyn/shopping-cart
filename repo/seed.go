package repo

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seedProducts(db *gorm.DB) {
	db.Create(&Product{
		Id:    1,
		Name:  "Product 1",
		Price: 1.99,
	})
	db.Create(&Product{
		Id:    2,
		Name:  "Product 2",
		Price: 2.99,
	})
	db.Create(&Product{
		Id:    3,
		Name:  "Product 3",
		Price: 3.99,
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func seedUsers(db *gorm.DB) {
	// for now, I don't care about the errors
	user1PasswordHash, _ := hashPassword("123456")
	db.Create(&User{
		Email:        "john@gmail.com",
		PasswordHash: user1PasswordHash,
	})

	user2PasswordHash, _ := hashPassword("password")
	db.Create(&User{
		Email:        "jane@gmail.com",
		PasswordHash: user2PasswordHash,
	})
}

func trunkateTables(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users RESTART IDENTITY")
	db.Exec("TRUNCATE TABLE products RESTART IDENTITY")

}

func SeedDb() {
	db := getDb()

	trunkateTables(db)

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})

	seedProducts(db)
	seedUsers(db)
}
