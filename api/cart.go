package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kolaczyn/shopping-cart/application"
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

	newCart := application.UpdateCart(cart)
	c.JSON(http.StatusOK, newCart)
}
