package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"log-service/internal/data"
	"log-service/logs"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	if err := l.Models.LogEntry.Insert(logEntry); err != nil {
		res := &logs.LogResponse{
			Result: "Failed",
		}

		return res, err
	}

	// return response
	res := &logs.LogResponse{
		Result: "logged!",
	}

	return res, nil
}

func (app *Config) gRPCListen() {
	listen, err := net.Listen("tcp", ":"+_gRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen gRPC: %v\n", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC Server started on port %s\n", _gRPCPort)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to listen gRPC: %v\n", err)
	}
}
