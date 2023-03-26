package repo

import "gorm.io/gorm"

type CartItem struct {
	Id       int
	Quantity int
}

type Cart struct {
	Items  []CartItem
	UserId string
}

type Product struct {
	gorm.Model
	Id    int
	Name  string
	Price float64
}

type User struct {
	gorm.Model
	Id           int
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
}
