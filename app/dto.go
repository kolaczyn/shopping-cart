package app

type CartItemDto struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type CartDto struct {
	Items []CartItemDto `json:"items"`
}
