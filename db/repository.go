package db

func CreateEmptyCart() Cart {
	items := []CartItem{}
	cart := Cart{
		Items: items,
	}
	return cart
}

func CreateCartWithDummyData() Cart {
	items := []CartItem{
		{
			Id:       1,
			Price:    10.0,
			Quantity: 1,
		},
		{
			Id:       2,
			Price:    20.0,
			Quantity: 2,
		},
	}
	cart := Cart{
		Items: items,
	}
	return cart
}
