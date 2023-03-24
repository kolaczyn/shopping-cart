package app

import "github.com/kolaczyn/shopping-cart/repo"

func (dto CartDto) dtoToDb() repo.CartDb {
	items := []repo.CartItemDb{}
	for _, item := range dto.Items {
		items = append(items, repo.CartItemDb{
			Id:       item.Id,
			Quantity: item.Quantity,
		})
	}
	cart := repo.CartDb{
		Items: items,
	}
	return cart
}

func dbToDto(db repo.CartDb) CartDto {
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
