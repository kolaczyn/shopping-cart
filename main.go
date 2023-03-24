package main

import "github.com/kolaczyn/shopping-cart/api"

type User struct {
	Username string
	Id       string
}

func main() {
	// seed database if this is the first time running this app
	// db.SeedDb()

	api.Run()
}
