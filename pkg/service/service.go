package service

import (
	authpb "conference/pb/auth"
	pb "conference/pb/conference"
	"conference/pkg/client/auth"
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
	Client      auth.AuthClient
	PrivateRepo interfaces.PrivateRepo
	GroupRepo   interfaces.GroupRepo
	PublicRepo  interfaces.PublicRepo
}

func NewConferenceServer(client auth.AuthClient, pRepo interfaces.PrivateRepo, gRepo interfaces.GroupRepo, puRepo interfaces.PublicRepo) *ConferenceServer {
	return &ConferenceServer{
		Client:      client,
		PrivateRepo: pRepo,
		GroupRepo:   gRepo,
		PublicRepo:  puRepo,
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

func (s *ConferenceServer) StartPrivateConference(ctx context.Context, req *pb.StartPrivateConferenceRequest) (*pb.StartPrivateConferenceResponse, error) {
	traceID := ctx.Value("traceID")
	var input utility.PrivateRoom
	copier.Copy(&input, req)
	uid, err := utility.UID(8)
	if err != nil {
		return nil, err
	}
	input.ConferenceID = uid
	_, err = s.PrivateRepo.CreatePrivateRoom(input)
	if err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	participantInput := utility.PrivateRoomParticipants{
		UserID:       req.UserID,
		ConferenceID: input.ConferenceID,
		Permission:   true,
		CamStatus:    "active",
		MicStatus:    "active",
		JoinTime:     time.Now(),
		Role:         "host",
	}
	if err = s.PrivateRepo.AddParticipantInPrivateRoom(participantInput); err != nil {
		log.Fatal(err, traceID)
		return nil, err
	}
	response := pb.StartPrivateConferenceResponse{
		ConferenceID: input.ConferenceID,
	}
	log.Println("conference room created")
	return &response, nil
}

func (s *ConferenceServer) StartGroupConference(ctx context.Context, req *pb.StartGroupConferenceRequest) (*pb.StartGroupConferenceResponse, error) {
	traceID := ctx.Value("traceID")
	var input utility.GroupRoom
	copier.Copy(&input, req)
	request := &authpb.GroupHostPermissionRequest{}
	_, err := s.Client.GroupHostPermission(ctx, request)
	if err != nil {
		return nil, err
	}
	// if permission == false {
	// 	return nil, err
	// }

	uid, err := utility.UID(8)
	if err != nil {
		return nil, err
	}
	input.ConferenceID = uid
	if err = s.GroupRepo.CreateGroupRoom(input); err != nil {
		return nil, err
	}
	participantInput := utility.GroupRoomParticipants{
		UserID:       req.UserID,
		GroupID:      input.GroupID,
		ConferenceID: input.ConferenceID,
		Permission:   true,
		CamStatus:    "active",
		MicStatus:    "active",
		JoinTime:     time.Now(),
		Role:         "host",
	}
	if err = s.GroupRepo.AddParticipantInGroupRoom(participantInput); err != nil {
		log.Fatal(err, traceID)
		return nil, err
	}
	response := pb.StartGroupConferenceResponse{
		Result:       "conference room created",
		ConferenceID: input.ConferenceID,
	}
	return &response, nil
}

func (s *ConferenceServer) StartPublicConference(ctx context.Context, req *pb.StartPublicConferenceRequest) (*pb.StartPublicConferenceResponse, error) {
	var input utility.PublicRoom
	traceID := ctx.Value("traceID")
	request := &authpb.PublicHostPermissionRequest{}
	_, err := s.Client.PublicHostPermission(ctx, request)
	if err != nil {
		return nil, err
	}
	// if permission == false {
	// 	return nil, err
	// }
	copier.Copy(&input, req)
	uid, err := utility.UID(8)
	if err != nil {
		return nil, err
	}
	input.ConferenceID = uid
	if err = s.PublicRepo.CreatePublicRoom(input); err != nil {
		return nil, err
	}
	participantInput := utility.PublicRoomParticipants{
		UserID:       req.UserID,
		ConferenceID: input.ConferenceID,
		Permission:   true,
		CamStatus:    "active",
		MicStatus:    "active",
		JoinTime:     time.Now(),
		Role:         "host",
	}
	if err = s.PublicRepo.AddParticipantInPublicRoom(participantInput); err != nil {
		log.Fatal(err, traceID)
		return nil, err
	}
	response := pb.StartPublicConferenceResponse{
		Result:       "conference room created",
		ConferenceID: input.ConferenceID,
	}
	return &response, nil
}

func (s *ConferenceServer) JoinPrivateConference(ctx context.Context, req *pb.JoinPrivateConferenceRequest) (*pb.JoinPrivateConferenceResponse, error) {
	conferenceID := req.ConferenceID
	userID := req.UserID
	response := pb.JoinPrivateConferenceResponse{}
	participantLimit, err := s.PrivateRepo.CheckLimit(conferenceID)
	if err != nil {
		log.Fatal("checklimit", err, ctx.Value("traceID"))
		return nil, err
	}
	currentParticipants, err := s.PrivateRepo.CountParticipants(conferenceID)
	if err != nil {
		log.Fatal("countparticipant", err, ctx.Value("traceID"))
		return nil, err
	}
	permission, err := s.PrivateRepo.CheckParticipantPermission(conferenceID, userID)
	if err != nil {
		log.Fatal("checkpermission", err, ctx.Value("traceID"))
		return nil, err
	}
	if permission == false {
		response = pb.JoinPrivateConferenceResponse{
			Result: "Participant permission denied",
		}
		return &response, errors.New("Participant permission denied")
	}
	if currentParticipants < participantLimit {
		response = pb.JoinPrivateConferenceResponse{
			Result: "Join request send",
		}
		return &response, nil
	} else {
		response = pb.JoinPrivateConferenceResponse{
			Result: "Participant limit exceeded",
		}
		return &response, errors.New("Participant limit exceeded")
	}
}

func (s *ConferenceServer) JoinGroupConference(ctx context.Context, req *pb.JoinGroupConferenceRequest) (*pb.JoinGroupConferenceResponse, error) {
	var input utility.GroupRoomParticipants
	traceID := ctx.Value("traceID")
	permission, err := s.Client.GroupParticipantPermission(input.UserID)
	if err != nil {
		return nil, errors.New("error" + err + "traceID" + traceID)
	}
	if permission == false {
		return nil, errors.New("error" + err + "traceID" + traceID)
	}
	participantInput := utility.GroupRoomParticipants{
		UserID:       req.UserID,
		GroupID:      input.GroupID,
		ConferenceID: input.ConferenceID,
		Permission:   true,
		CamStatus:    "active",
		MicStatus:    "active",
		JoinTime:     time.Now(),
		Role:         input.Role,
	}
	if err = s.GroupRepo.AddParticipantInGroupRoom(participantInput); err != nil {
		log.Fatal(err, traceID)
		return nil, err
	}
	response := pb.JoinGroupConferenceResponse{
		Result: "Joined group conference",
	}
	return &response, nil
}

func (s *ConferenceServer) JoinPublicConference(ctx context.Context, req *pb.JoinPublicConferenceRequest) (*pb.JoinPublicConferenceResponse, error) {
	var input utility.PublicRoomParticipants
	traceID := ctx.Value("traceID")
	permission, err := s.Client.PublicParticipantPermission(input.UserID)
	if err != nil {
		return nil, errors.New("error" + err + "traceID" + traceID)
	}
	if permission == false {
		return nil, errors.New("error" + err + "traceID" + traceID)
	}
	participantInput := utility.PublicRoomParticipants{
		UserID:       req.UserID,
		ConferenceID: input.ConferenceID,
		Permission:   true,
		CamStatus:    "active",
		MicStatus:    "active",
		JoinTime:     time.Now(),
		Role:         input.Role,
	}
	if err = s.PublicRepo.AddParticipantInPublicRoom(participantInput); err != nil {
		log.Fatal(err, traceID)
		return nil, err
	}
	response := pb.JoinPublicConferenceResponse{
		Result: "joined public conference",
	}
	return &response, nil
}

func (s *ConferenceServer) AcceptJoining(ctx context.Context, req *pb.AcceptJoiningRequest) (*pb.AcceptJoiningResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	participantInput := utility.PrivateRoomParticipants{
		UserID:       userID,
		ConferenceID: conferenceID,
		Permission:   true,
		CamStatus:    "active",
		MicStatus:    "active",
		JoinTime:     time.Now(),
		Role:         "user",
	}

	err = s.PrivateRepo.AddParticipantInPrivateRoom(participantInput)
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
		if err := s.PrivateRepo.BlockParticipant(conferenceID, userID); err != nil {
			log.Fatal(err, ctx.Value("traceID"))
			return nil, err
		}
	}
	if err = s.PrivateRepo.RemoveParticipant(conferenceID, userID); err != nil {
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
	participantInput := utility.PrivateRoomParticipants{
		UserID:       userID,
		ConferenceID: conferenceID,
		ExitTime:     time.Now(),
	}
	if err = s.PrivateRepo.UpdateParticipantExitTime(participantInput); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	if err = s.PrivateRepo.RemoveParticipant(conferenceID, userID); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	response := pb.LeaveConferenceResponse{
		Result: "Exited from the conference",
	}
	return &response, nil
}

func (s *ConferenceServer) EndConference(ctx context.Context, req *pb.EndConferenceRequest) (*pb.EndConferenceResponse, error) {
	err = s.PrivateRepo.RemoveRoom(req.ConferenceID)
	if err != nil {
		return nil, err
	}
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
