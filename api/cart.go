package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appCart "github.com/kolaczyn/shopping-cart/app/cart"
)

func addCartRoutes(router *gin.RouterGroup) {
	router.GET("/", getCart)
	router.POST("/", postCart)
	router.DELETE("/", deleteCart)

	router.GET("/summary", getCartSummary)
}

func getCart(c *gin.Context) {
	cart, err := appCart.GetCart()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, cart)
}

func postCart(c *gin.Context) {
	var cart appCart.CartDto
	c.BindJSON(&cart)

	newCart, err := appCart.UpdateCart(cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, newCart)
}

func deleteCart(c *gin.Context) {
	err := appCart.DeleteCart()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func getCartSummary(c *gin.Context) {
	summary, err := appCart.GetCartSummary()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, summary)
}
