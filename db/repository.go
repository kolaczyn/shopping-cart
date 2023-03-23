package db

var cartDatabase CartDb = initCart()

func initCart() CartDb {
	items := []CartItemDb{}
	cart := CartDb{
		Items: items,
	}
	return cart
}

func UpdateCart(cart CartDb) CartDb {
	cartDatabase.Items = cart.Items

	return cartDatabase
}

func GetCart() CartDb {
	return cartDatabase
}
