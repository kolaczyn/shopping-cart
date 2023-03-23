package application

import "github.com/kolaczyn/shopping-cart/db"

func UpdateCart(cart CartDto) CartDto {
	newCart := db.UpdateCart(dtoToDb(cart))
	return dbToDto(newCart)
}
