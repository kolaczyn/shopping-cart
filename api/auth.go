package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appAuth "github.com/kolaczyn/shopping-cart/app/auth"
)

func addAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", login)
	router.POST("/register", register)
}

func register(c *gin.Context) {
	var form appAuth.RegisterFormDto
	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := appAuth.Register(form.Email, form.Password)
	if err != nil {
		// TODO this shouldn't always return 400
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

func login(c *gin.Context) {
	var form appAuth.LoginFormDto
	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := appAuth.Login(form.Email, form.Password)
	if err != nil {
		// TODO this shouldn't always return 400
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
