package appCart

import (
	db "github.com/kolaczyn/shopping-cart/repo"
)

func UpdateCart(cart CartDto) (CartDto, error) {
	err := CheckIsCartValid(cart)
	if err != nil {
		return CartDto{}, err
	}
	newCart, err := db.UpdateCart(*cart.dtoToDb())
	return *dbToDto(newCart), err
}
