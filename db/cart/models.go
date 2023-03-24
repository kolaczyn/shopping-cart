package db

type CartItemDb struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type CartDb struct {
	Items  []CartItemDb `json:"items"`
	UserId string       `json:"userId"`
}
