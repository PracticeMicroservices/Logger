package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() (*mongo.Client, error) {
	//create connection options
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	//connect to mongo
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting: ", err)
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return c, nil
}
