package service

import (
	"conference/pb"
	"context"
	"log"
)

type Conference struct {
	pb.UnimplementedConferenceServer
}

func (s *Conference) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Conference: Health Checked")

	result := "Hello, " + req.Data
	return &pb.Response{Result: result}, nil
}

func (s *Conference) StartConference(ctx context.Context, req *pb.StartConferenceRequest) (*pb.StartConferenceResponse, error) {

	response := pb.StartConferenceResponse{
		ConferenceID: "rean101",
	}
	return &response, nil
}
