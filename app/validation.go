package app

import (
	"errors"

	"github.com/kolaczyn/shopping-cart/repo"
	"github.com/samber/lo"
)

func CheckIsCartValid(cart CartDto) error {
	products := repo.GetAllProducts()

	err := checkAllProductsExisting(cart.Items, products)
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

	return nil
}

func checkAllProductsExisting(cartItems []CartItemDto, products []repo.ProductDb) error {
	for _, item := range cartItems {
		_, isFound := lo.Find(products, func(product repo.ProductDb) bool {
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
	var result []int = lo.Map(items, func(item CartItemDto, _ int) int {
		return item.Id
	})
	// I can probably make it shorter by using lo.FindDuplicatesBy
	duplicates := lo.FindDuplicates(result)
	if len(duplicates) > 0 {
		return errors.New("there are duplicates")
	}
	return nil
}
