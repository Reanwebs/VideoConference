package main

import (
	"conference/pb"
	"conference/pkg/service"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	ConferenceService := &service.Conference{}

	pb.RegisterConferenceServer(grpcServer, ConferenceService)

	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Conference service is running...5050")
	go grpcServer.Serve(listener)

	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatalf("Failed to start health check server: %v", err)
	}
}
