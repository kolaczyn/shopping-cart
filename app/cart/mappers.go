package appCart

import "github.com/kolaczyn/shopping-cart/repo"

func (dto *CartDto) dtoToDb() *repo.Cart {
	items := []repo.CartItem{}
	for _, item := range dto.Items {
		items = append(items, repo.CartItem{
			Id:       item.Id,
			Quantity: item.Quantity,
		})
	}
	cart := repo.Cart{
		Items: items,
	}
	return &cart
}

func dbToDto(db *repo.Cart) *CartDto {
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
	return &cart
}
