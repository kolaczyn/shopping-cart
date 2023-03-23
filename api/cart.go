package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	application "github.com/kolaczyn/shopping-cart/application/cart"
)

func addCartRoutes(router *gin.RouterGroup) {
	router.GET("/", getCart)
	router.POST("/", postCart)
}

func getCart(c *gin.Context) {
	cart := application.GetCart()
	c.JSON(http.StatusOK, cart)
}

func postCart(c *gin.Context) {
	var cart application.CartDto
	c.BindJSON(&cart)

	newCart, err := application.UpdateCart(cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, newCart)
}
