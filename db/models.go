package db

type CartItem struct {
	Id       int     `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}

func (cart *Cart) TotalCost() float64 {
	var total float64
	for _, item := range cart.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
