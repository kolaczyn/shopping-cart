package appCart

import "github.com/kolaczyn/shopping-cart/repo"

func GetCart() (*CartDto, error) {
	dbCart, err := repo.GetCart()
	return dbToDto(dbCart), err
}
