package application

import db "github.com/kolaczyn/shopping-cart/db/cart"

func (dto CartDto) dtoToDb() db.CartDb {
	items := []db.CartItemDb{}
	for _, item := range dto.Items {
		items = append(items, db.CartItemDb{
			Id:       item.Id,
			Quantity: item.Quantity,
		})
	}
	cart := db.CartDb{
		Items: items,
	}
	return cart
}

func dbToDto(db db.CartDb) CartDto {
	items := []CartItemDto{}
	for _, item := range db.Items {
		items = append(items, CartItemDto{
			Id:       item.Id,
			Quantity: item.Quantity,
		})
	}
	cart := CartDto{
		Items: items,
	}
	return cart
}
