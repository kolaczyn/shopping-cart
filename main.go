package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartItem struct {
	Id       int     `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}

func createEmptyCart() Cart {
	items := []CartItem{}
	cart := Cart{
		Items: items,
	}
	return cart
}

func createCartWithDummyData() Cart {
	items := []CartItem{
		{
			Id:       1,
			Price:    10.0,
			Quantity: 1,
		},
		{
			Id:       2,
			Price:    20.0,
			Quantity: 2,
		},
	}
	cart := Cart{
		Items: items,
	}
	return cart
}

func (cart *Cart) totalCost() float64 {
	var total float64
	for _, item := range cart.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func main() {
	r := gin.Default()
	cart := createCartWithDummyData()

	r.GET("v1/cart", func(c *gin.Context) {
		c.JSON(http.StatusOK, cart)
	})

	r.GET("/v1/cart/total", func(c *gin.Context) {
		c.JSON(http.StatusOK, cart.totalCost())
	})
	r.Run()
}
