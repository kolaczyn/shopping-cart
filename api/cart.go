package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kolaczyn/shopping-cart/application"
)

func addCartRoutes(router *gin.RouterGroup) {

	router.GET("/cart", func(c *gin.Context) {
		cart := application.GetCart()
		c.JSON(http.StatusOK, cart)
	})

	router.POST("/cart", func(c *gin.Context) {
		// TODO read cart from request
		cart := application.CartDto{
			Items: []application.CartItemDto{
				{
					Id:       123,
					Quantity: 1,
				},
				{
					Id:       456,
					Quantity: 2,
				},
			},
		}

		newCart := application.UpdateCart(cart)
		c.JSON(http.StatusOK, newCart)
	})

}
