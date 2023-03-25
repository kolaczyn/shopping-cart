package app

import (
	"github.com/kolaczyn/shopping-cart/repo"
	"github.com/samber/lo"
)

func getProductIdsDto(items []CartItemDto) []int {
	return lo.Map(items, func(item CartItemDto, _ int) int {
		return item.Id
	})
}

func getProductIdsDb(items []repo.CartItemDb) []int {
	return lo.Map(items, func(item repo.CartItemDb, _ int) int {
		return item.Id
	})
}
