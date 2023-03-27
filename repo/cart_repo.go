package repo

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userId = "2137"

func getCartCollection() *mongo.Collection {
	uri := os.Getenv("SHOPPING_CART_DB_URL")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		// It doesn't make sense to continue the program if we can't connect to the database
		panic(err)
	}

	collection := client.Database("test").Collection("cart")
	return collection
}

func UpdateCart(cart Cart) (*Cart, error) {
	opts := options.Update().SetUpsert(true)
	// TODO check if the cart exists
	_, err := getCartCollection().UpdateOne(context.Background(), bson.M{"userId": userId}, bson.M{"$set": cart}, opts)
	return &cart, err
}

func GetCart() (*Cart, error) {
	var cart = Cart{}
	err := getCartCollection().FindOne(context.Background(), bson.M{"userId": userId}).Decode(&cart)
	return &cart, err
}

func DeleteCart() error {
	// TODO check if the cart exists
	_, err := getCartCollection().DeleteOne(context.Background(), bson.M{"userId": userId})
	return err
}
