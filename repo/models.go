package repo

import "gorm.io/gorm"

type CartItemDb struct {
	Id       int
	Quantity int
}

type CartDb struct {
	Items  []CartItemDb
	UserId string
}

type ProductDb struct {
	gorm.Model
	Id    int
	Name  string
	Price float64
}
