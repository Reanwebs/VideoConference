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
		ConferenceID: 101,
	}
	return &response, nil
}

func (s *Conference) JoinConference(ctx context.Context, req *pb.JoinConferenceRequest) (*pb.JoinConferenceResponse, error) {
	response := pb.JoinConferenceResponse{
		Result: "Join request send",
	}
	return &response, nil
}

func (s *Conference) AcceptJoining(ctx context.Context, req *pb.AcceptJoiningRequest) (*pb.AcceptJoiningResponse, error) {
	response := pb.AcceptJoiningResponse{
		Result: "Join request accepted",
	}
	return &response, nil
}

func (s *Conference) DeclineJoining(ctx context.Context, req *pb.DeclineJoiningRequest) (*pb.DeclineJoiningResponse, error) {
	response := pb.DeclineJoiningResponse{
		Result: "Join request declined",
	}
	return &response, nil
}

func (s *Conference) RemoveParticipant(ctx context.Context, req *pb.RemoveParticipantRequest) (*pb.RemoveParticipantResponse, error) {
	response := pb.RemoveParticipantResponse{
		Result: "participant removed",
	}
	return &response, nil
}

func (s *Conference) ToggleCamera(ctx context.Context, req *pb.ToggleCameraRequest) (*pb.ToggleCameraResponse, error) {
	response := pb.ToggleCameraResponse{
		Result: "Camera toggled",
	}
	return &response, nil
}

func (s *Conference) ToggleMic(ctx context.Context, req *pb.ToggleMicRequest) (*pb.ToggleMicResponse, error) {
	response := pb.ToggleMicResponse{
		Result: "Mic toggled",
	}
	return &response, nil
}

func (s *Conference) ToggleParticipantCamera(ctx context.Context, req *pb.ToggleParticipantCameraRequest) (*pb.ToggleParticipantCameraResponse, error) {
	response := pb.ToggleParticipantCameraResponse{
		Result: "Participant camera toggled",
	}
	return &response, nil
}

func (s *Conference) ToggleParticipantMic(ctx context.Context, req *pb.ToggleParticipantMicRequest) (*pb.ToggleParticipantMicResponse, error) {
	response := pb.ToggleParticipantMicResponse{
		Result: "Participant mic toggled",
	}
	return &response, nil
}

func (s *Conference) EndConference(ctx context.Context, req *pb.EndConferenceRequest) (*pb.EndConferenceResponse, error) {
	response := pb.EndConferenceResponse{
		Result: "Conference ended",
	}
	return &response, nil
}
