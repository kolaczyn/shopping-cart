package appCart

import (
	"github.com/kolaczyn/shopping-cart/repo"
	"github.com/samber/lo"
)

func GetCartSummary() (*CartSummaryDto, error) {
	dbCart, err := repo.GetCart()
	if err != nil {
		return nil, err
	}
	productIds := getProductIdsDb(dbCart.Items)
	products, err := repo.GetProductsByIds(productIds)
	if err != nil {
		return nil, err
	}

	cartSummary, err := mergeProductsWithCart(dbCart.Items, *products)
	return cartSummary, err
}

// TODO add tests
func mergeProductsWithCart(cartItems []repo.CartItem, products []repo.Product) (*CartSummaryDto, error) {
	cartSummaryItems := lo.FilterMap(cartItems, func(cartItem repo.CartItem, _ int) (CartItemSummaryDto, bool) {
		// yeah, I know thia is something like O(n^2) but I don't care :p
		product, isFound := lo.Find(products, func(product repo.Product) bool {
			return product.Id == cartItem.Id
		})
		if !isFound {
			return CartItemSummaryDto{}, false
		}

		summaryItem := CartItemSummaryDto{
			Id:       product.Id,
			Quantity: cartItem.Quantity,
			Price:    product.Price,
		}

		return summaryItem, true

	})

	roundedTotal := lo.Reduce(cartSummaryItems, func(total float64, item CartItemSummaryDto, _ int) float64 {
		return total + item.Price*float64(item.Quantity)
	}, 0)

	roundedTotal = float64(int(roundedTotal*100)) / 100

	return &CartSummaryDto{
		Items: cartSummaryItems,
		Total: roundedTotal,
	}, nil
}
