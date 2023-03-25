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

type RegisterFormDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginFormDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
