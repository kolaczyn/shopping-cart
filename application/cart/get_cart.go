package application

import db "github.com/kolaczyn/shopping-cart/db/cart"

func GetCart() CartDto {
	dbCart := db.GetCart()
	return dbToDto(dbCart)
}
