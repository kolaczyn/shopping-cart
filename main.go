package main

import "github.com/kolaczyn/shopping-cart/api"

type User struct {
	Username string
	Id       string
}

func main() {
	api.Run()
}
