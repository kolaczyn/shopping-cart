package api

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run() {
	getRoutes()
	router.Run(":8080")
}

func getRoutes() {
	cartRoutes := router.Group("/v1/cart")
	addCartRoutes(cartRoutes)
}
