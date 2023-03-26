package main

import (
	"github.com/joho/godotenv"
	"github.com/kolaczyn/shopping-cart/api"
)

type User struct {
	Username string
	Id       string
}

func main() {
	godotenv.Load()
	// seed database if this is the first time running this app
	// repo.SeedDb()

	api.Run()
}
