package service

import (
	"conference/pb"
	"conference/pkg/common/utility"
	"conference/pkg/repository/interfaces"
	"context"
	"log"

	"github.com/jinzhu/copier"
)

type ConferenceServer struct {
	pb.UnimplementedConferenceServer
	Repo interfaces.ConferenceRepo
}

func NewConferenceServer(repo interfaces.ConferenceRepo) *ConferenceServer {
	return &ConferenceServer{
		Repo: repo,
	}
}

func (s *ConferenceServer) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Conference: Health Checked")
	result := "Hello, " + req.Data
	return &pb.Response{Result: result}, nil
}

func (s *ConferenceServer) StartConference(ctx context.Context, req *pb.StartConferenceRequest) (*pb.StartConferenceResponse, error) {
	var input utility.ConferenceRoom
	copier.Copy(&input, req)
	conferenceID, err := s.Repo.CreateRoom(input)
	if err != nil {
		log.Fatal(err)
	}
	response := pb.StartConferenceResponse{
		ConferenceID: int32(conferenceID),
	}
	return &response, nil
}

func (s *ConferenceServer) JoinConference(ctx context.Context, req *pb.JoinConferenceRequest) (*pb.JoinConferenceResponse, error) {
	response := pb.JoinConferenceResponse{
		Result: "Join request send",
	}
	return &response, nil
}

func (s *ConferenceServer) AcceptJoining(ctx context.Context, req *pb.AcceptJoiningRequest) (*pb.AcceptJoiningResponse, error) {
	response := pb.AcceptJoiningResponse{
		Result: "Join request accepted",
	}
	return &response, nil
}

func (s *ConferenceServer) DeclineJoining(ctx context.Context, req *pb.DeclineJoiningRequest) (*pb.DeclineJoiningResponse, error) {
	response := pb.DeclineJoiningResponse{
		Result: "Join request declined",
	}
	return &response, nil
}

func (s *ConferenceServer) RemoveParticipant(ctx context.Context, req *pb.RemoveParticipantRequest) (*pb.RemoveParticipantResponse, error) {
	response := pb.RemoveParticipantResponse{
		Result: "participant removed",
	}
	return &response, nil
}

func (s *ConferenceServer) ToggleCamera(ctx context.Context, req *pb.ToggleCameraRequest) (*pb.ToggleCameraResponse, error) {
	response := pb.ToggleCameraResponse{
		Result: "Camera toggled",
	}
	return &response, nil
}

func (s *ConferenceServer) ToggleMic(ctx context.Context, req *pb.ToggleMicRequest) (*pb.ToggleMicResponse, error) {
	response := pb.ToggleMicResponse{
		Result: "Mic toggled",
	}
	return &response, nil
}

func (s *ConferenceServer) ToggleParticipantCamera(ctx context.Context, req *pb.ToggleParticipantCameraRequest) (*pb.ToggleParticipantCameraResponse, error) {
	response := pb.ToggleParticipantCameraResponse{
		Result: "Participant camera toggled",
	}
	return &response, nil
}

func (s *ConferenceServer) ToggleParticipantMic(ctx context.Context, req *pb.ToggleParticipantMicRequest) (*pb.ToggleParticipantMicResponse, error) {
	response := pb.ToggleParticipantMicResponse{
		Result: "Participant mic toggled",
	}
	return &response, nil
}

func (s *ConferenceServer) EndConference(ctx context.Context, req *pb.EndConferenceRequest) (*pb.EndConferenceResponse, error) {
	response := pb.EndConferenceResponse{
		Result: "Conference ended",
	}
	return &response, nil
}

func (s *ConferenceServer) mustEmbedUnimplementedConferenceServer() {

}
