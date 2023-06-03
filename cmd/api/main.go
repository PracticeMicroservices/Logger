package main

import (
	"context"
	"log"
	"logger/database"
	"time"
)

func main() {
	//connect to mongo
	mongoClient, err := database.ConnectToMongo()
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := NewApp(mongoClient)

	log.Println("Starting Logger service on port 80")

	//start server
	app.serve()
}
