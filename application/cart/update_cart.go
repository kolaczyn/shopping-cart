package application

import db "github.com/kolaczyn/shopping-cart/db/cart"

func UpdateCart(cart CartDto) CartDto {
	newCart := db.UpdateCart(cart.dtoToDb())
	return dbToDto(newCart)
}
