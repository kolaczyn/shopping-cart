package db

func GetAllProducts() []ProductDb {
	return []ProductDb{
		{
			Id:    1,
			Name:  "Product 1",
			Price: 1.99,
		},
		{
			Id:    2,
			Name:  "Product 2",
			Price: 2.99,
		},
		{
			Id:    3,
			Name:  "Product 3",
			Price: 3.99,
		},
	}
}
