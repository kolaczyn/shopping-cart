package application

import db "github.com/kolaczyn/shopping-cart/db/cart"

func DeleteCart() error {
	err := db.DeleteCart()
	return err
}
