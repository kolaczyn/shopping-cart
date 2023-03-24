package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userId = "2137"

func getCartCollection() *mongo.Collection {
	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		// It doesn't make sense to continue the program if we can't connect to the database
		panic(err)
	}

	collection := client.Database("test").Collection("cart")
	return collection
}

func UpdateCart(cart CartDb) (CartDb, error) {
	opts := options.Update().SetUpsert(true)
	// TODO check if the cart exists
	result, err := getCartCollection().UpdateOne(context.Background(), bson.M{"userId": userId}, bson.M{"$set": cart}, opts)
	println("Updated", result.ModifiedCount, "documents")
	return cart, err
}

func GetCart() (CartDb, error) {
	var cart = CartDb{}
	err := getCartCollection().FindOne(context.Background(), bson.M{"userId": userId}).Decode(&cart)
	return cart, err
}

func DeleteCart() error {
	// TODO check if the cart exists
	_, err := getCartCollection().DeleteOne(context.Background(), bson.M{"userId": userId})
	return err
}
