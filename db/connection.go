package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN is the object with the function of connection*/
var MongoCN = ConnectionDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://dach3r:Password.1090@cluster0.jdqyf.mongodb.net/test")

/*ConnectionDB if use for connect with mongo*/
func ConnectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Success connection with DB")
	return client
}

/*VerifyConnection if used for check if connection is success*/
func VerifyConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}

	return true
}
