package app

import "github.com/kolaczyn/shopping-cart/repo"

func DeleteCart() error {
	err := repo.DeleteCart()
	return err
}
