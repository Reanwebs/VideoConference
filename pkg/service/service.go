package service

import (
	"conference/pb"
	"conference/pkg/common/utility"
	"conference/pkg/repository/interfaces"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jinzhu/copier"
)

var (
	err error
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
	timestamp := time.Now().UnixNano()
	randomNumber := rand.Intn(1000000)
	traceID := fmt.Sprintf("%d-%d", timestamp, randomNumber)
	result := traceID + " " + req.Data
	return &pb.Response{Result: result}, nil
}

func (s *ConferenceServer) StartConference(ctx context.Context, req *pb.StartConferenceRequest) (*pb.StartConferenceResponse, error) {
	var input utility.ConferenceRoom
	copier.Copy(&input, req)
	uid := utility.UID(8)
	input.ConferenceID = uid
	_, err := s.Repo.CreateRoom(input)
	if err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	participantInput := utility.ConferenceParticipants{
		UserID:       req.UserID,
		ConferenceID: input.ConferenceID,
		Permission:   true,
		CamStatus:    "active",
		MicStatus:    "active",
		JoinTime:     time.Now(),
		Role:         "host",
	}
	if err = s.Repo.AddParticipant(participantInput); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	response := pb.StartConferenceResponse{
		ConferenceID: input.ConferenceID,
	}
	log.Println("conference room created")
	return &response, nil
}

func (s *ConferenceServer) JoinConference(ctx context.Context, req *pb.JoinConferenceRequest) (*pb.JoinConferenceResponse, error) {
	conferenceID := req.ConferenceID
	userID := req.UserID
	response := pb.JoinConferenceResponse{}
	participantLimit, err := s.Repo.CheckLimit(conferenceID)
	if err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	currentParticipants, err := s.Repo.CountParticipants(conferenceID)
	if err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	permission, err := s.Repo.CheckParticipantPermission(conferenceID, userID)
	if err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	if permission == false {
		response = pb.JoinConferenceResponse{
			Result: "Participant permission denied",
		}
		return &response, errors.New("Participant permission denied")
	}
	if currentParticipants < participantLimit {
		response = pb.JoinConferenceResponse{
			Result: "Join request send",
		}
		return &response, nil
	} else {
		response = pb.JoinConferenceResponse{
			Result: "Participant limit exceeded",
		}
		return &response, errors.New("Participant limit exceeded")
	}
}

func (s *ConferenceServer) AcceptJoining(ctx context.Context, req *pb.AcceptJoiningRequest) (*pb.AcceptJoiningResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	participantInput := utility.ConferenceParticipants{
		UserID:       userID,
		ConferenceID: conferenceID,
		Permission:   true,
		CamStatus:    "active",
		MicStatus:    "active",
		JoinTime:     time.Now(),
		Role:         "user",
	}

	err = s.Repo.AddParticipant(participantInput)
	if err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	} else {
		response := pb.AcceptJoiningResponse{
			Result: "Join request accepted",
		}
		return &response, nil
	}
}

func (s *ConferenceServer) DeclineJoining(ctx context.Context, req *pb.DeclineJoiningRequest) (*pb.DeclineJoiningResponse, error) {
	response := pb.DeclineJoiningResponse{
		Result: "Join request declined",
	}
	return &response, nil
}

func (s *ConferenceServer) RemoveParticipant(ctx context.Context, req *pb.RemoveParticipantRequest) (*pb.RemoveParticipantResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	if req.Block == true {
		if err := s.Repo.BlockParticipant(conferenceID, userID); err != nil {
			log.Fatal(err, ctx.Value("traceID"))
			return nil, err
		}
	}
	if err = s.Repo.RemoveParticipant(conferenceID, userID); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	response := pb.RemoveParticipantResponse{
		Result: "participant removed",
	}
	return &response, nil
}

func (s *ConferenceServer) LeaveConference(ctx context.Context, req *pb.LeaveConferenceRequest) (*pb.LeaveConferenceResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	participantInput := utility.ConferenceParticipants{
		UserID:       userID,
		ConferenceID: conferenceID,
		ExitTime:     time.Now(),
	}
	if err = s.Repo.UpdateParticipantExitTime(participantInput); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	if err = s.Repo.RemoveParticipant(conferenceID, userID); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	response := pb.LeaveConferenceResponse{
		Result: "Participant mic toggled",
	}
	return &response, nil
}

func (s *ConferenceServer) EndConference(ctx context.Context, req *pb.EndConferenceRequest) (*pb.EndConferenceResponse, error) {
	err = s.Repo.RemoveRoom(req.ConferenceID)

	response := pb.EndConferenceResponse{
		Result: "Conference ended",
	}
	return &response, nil
}

// not implimented

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
