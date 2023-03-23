package application

import "github.com/kolaczyn/shopping-cart/db"

func GetCart() CartDto {
	dbCart := db.GetCart()
	return dbToDto(dbCart)
}
