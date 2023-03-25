package appCart

import "github.com/kolaczyn/shopping-cart/repo"

func DeleteCart() error {
	err := repo.DeleteCart()
	return err
}
