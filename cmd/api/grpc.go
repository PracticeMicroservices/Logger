package main

import (
	"context"
	"fmt"
	"log"
	"logger/data/models"
	"logger/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models models.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	fmt.Println("WriteLog")
	input := req.GetLogEntry()

	//write log
	logEntry := models.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	//return response
	res := &logs.LogResponse{Result: "success"}
	return res, nil
}

func (app *App) gGRPCListen() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal("Failed to listen on port 50001:", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{
		Models: app.Models,
	})

	log.Println("gRPC server listening on port 50001")

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC server over port 50001:", err)
	}
}
