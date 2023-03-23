package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kolaczyn/shopping-cart/db"
)

var cart = db.CreateCartWithDummyData()

func addCartRoutes(router *gin.RouterGroup) {

	router.GET("/cart", func(c *gin.Context) {
		c.JSON(http.StatusOK, cart)
	})

	router.GET("/cart/total", func(c *gin.Context) {
		c.JSON(http.StatusOK, cart.TotalCost())
	})
}
