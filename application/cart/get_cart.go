package application

import db "github.com/kolaczyn/shopping-cart/db/cart"

func GetCart() (CartDto, error) {
	dbCart, err := db.GetCart()
	return dbToDto(dbCart), err
}
