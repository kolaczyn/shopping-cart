package application

import (
	db "github.com/kolaczyn/shopping-cart/db/product"
	"github.com/samber/lo"
)

// TODO add tests checking:
// - if no duplicate productIds
// - if are products are existing
// - if quantities are positive
func CheckIsCartValid(cart CartDto) bool {
	return areAllProductsExisting(cart)
}

func areAllProductsExisting(cart CartDto) bool {
	products := db.GetAllProducts()

	for _, item := range cart.Items {
		_, isFound := lo.Find(products, func(product db.ProductDb) bool {
			return product.Id == item.Id
		})
		if !isFound {
			return false
		}
	}
	return true
}
