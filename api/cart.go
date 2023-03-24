package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	application "github.com/kolaczyn/shopping-cart/application/cart"
)

func addCartRoutes(router *gin.RouterGroup) {
	router.GET("/", getCart)
	router.POST("/", postCart)
	router.DELETE("/", deleteCart)
}

func getCart(c *gin.Context) {
	cart, err := application.GetCart()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
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

func deleteCart(c *gin.Context) {
	err := application.DeleteCart()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}
