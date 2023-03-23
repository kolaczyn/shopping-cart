package application

import (
	"errors"

	db "github.com/kolaczyn/shopping-cart/db/cart"
)

func UpdateCart(cart CartDto) (CartDto, error) {
	isValid := CheckIsCartValid(cart)
	if !isValid {
		return CartDto{}, errors.New("cart is not valid")
	}
	newCart := db.UpdateCart(cart.dtoToDb())
	return dbToDto(newCart), nil
}
