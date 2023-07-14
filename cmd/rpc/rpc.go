package rpc

import (
	"context"
	"fmt"
	"log"
	"logger/data/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Server var client *mongo.Client
type Server struct {
	mongoClient *mongo.Client
}

type Payload struct {
	Name string
	Data string
}

func NewServer(db *mongo.Client) *Server {
	return &Server{
		mongoClient: db,
	}
}

func (s *Server) LogInfo(payload Payload, reply *string) error {
	fmt.Println("Received payload via RPC:", payload.Name, payload.Data)
	collection := s.mongoClient.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), models.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into logs:", err)
		return err
	}

	*reply = "Processed payload via RPC: " + payload.Name + " " + payload.Data
	return nil
}
