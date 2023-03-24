package db

import "gorm.io/gorm"

type ProductDb struct {
	gorm.Model
	Id    int
	Name  string
	Price float64
}
