package appCart

import (
	"errors"

	"github.com/kolaczyn/shopping-cart/repo"
	"github.com/samber/lo"
)

func CheckIsCartValid(cart CartDto) error {
	productIds := getProductIdsDto(cart.Items)
	products, err := repo.GetProductsByIds(productIds)

	if err != nil {
		return err
	}

	err = checkPositiveQuantities(cart.Items)
	if err != nil {
		return err
	}

	err = checkForDuplicates(cart.Items)
	if err != nil {
		return err
	}

	err = checkAllProductsExisting(cart.Items, products)
	if err != nil {
		return err
	}

	return nil
}

func checkAllProductsExisting(cartItems []CartItemDto, products []repo.Product) error {
	for _, item := range cartItems {
		_, isFound := lo.Find(products, func(product repo.Product) bool {
			return product.Id == item.Id
		})
		if !isFound {
			return errors.New("not all products are existing")
		}
	}
	return nil
}

func checkPositiveQuantities(items []CartItemDto) error {
	for _, item := range items {
		if item.Quantity <= 0 {
			return errors.New("not all quantities are positive")
		}
	}
	return nil
}

func checkForDuplicates(items []CartItemDto) error {
	productIds := getProductIdsDto(items)
	duplicates := lo.FindDuplicates(productIds)
	if len(duplicates) > 0 {
		return errors.New("there are duplicates")
	}
	return nil
}
