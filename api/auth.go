package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kolaczyn/shopping-cart/app"
)

func addAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", login)
	router.POST("/register", register)
}

func register(c *gin.Context) {
	var form app.RegisterFormDto
	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := app.Register(form.Email, form.Password)
	if err != nil {
		// TODO this shouldn't always return 400
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

func login(c *gin.Context) {
	var form app.LoginFormDto
	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := app.Register(form.Email, form.Password)
	if err != nil {
		// TODO this shouldn't always return 400
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}
