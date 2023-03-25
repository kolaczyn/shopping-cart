package app

type CartItemDto struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type CartDto struct {
	Items []CartItemDto `json:"items"`
}

type CartItemSummaryDto struct {
	Id       int     `json:"id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type CartSummaryDto struct {
	Items []CartItemSummaryDto `json:"items"`
	Total float64              `json:"total"`
}
