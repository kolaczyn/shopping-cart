package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kolaczyn/shopping-cart/app"
)

func addCartRoutes(router *gin.RouterGroup) {
	router.GET("/", getCart)
	router.POST("/", postCart)
	router.DELETE("/", deleteCart)
}

func getCart(c *gin.Context) {
	cart, err := app.GetCart()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, cart)
}

func postCart(c *gin.Context) {
	var cart app.CartDto
	c.BindJSON(&cart)

	newCart, err := app.UpdateCart(cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, newCart)
}

func deleteCart(c *gin.Context) {
	err := app.DeleteCart()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}
