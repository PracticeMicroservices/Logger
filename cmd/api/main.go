package main

import (
	"context"
	"log"
	rpcServer "logger/cmd/rpc"
	"logger/database"
	"net"
	"net/rpc"
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
	RpcModel := rpcServer.NewServer(mongoClient)

	log.Println("Starting Logger service on port 80")

	err = rpc.Register(RpcModel)
	go app.rpcListen()

	go app.gGRPCListen()
	//start server
	app.serve()
}

func (a *App) rpcListen() {
	log.Println("Starting RPC server on port 5001")
	listen, err := net.Listen("tcp", "0.0.0.0:5001")

	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
