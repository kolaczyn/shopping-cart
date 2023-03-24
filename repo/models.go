package repo

import "gorm.io/gorm"

type CartItemDb struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type CartDb struct {
	Items  []CartItemDb `json:"items"`
	UserId string       `json:"userId"`
}

type ProductDb struct {
	gorm.Model
	Id    int
	Name  string
	Price float64
}
