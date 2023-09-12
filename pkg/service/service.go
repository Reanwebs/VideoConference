package service

import (
	authpb "conference/pb/auth"
	pb "conference/pb/conference"
	monitPb "conference/pb/monitization"
	"conference/pkg/client/auth"
	monit "conference/pkg/client/monitization"
	"conference/pkg/common/utility"
	"conference/pkg/repository/interfaces"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	err error
)

type ConferenceServer struct {
	pb.UnimplementedConferenceServer
	AuthClient  auth.AuthClient
	MonitClient monit.MonitizationClient
	PrivateRepo interfaces.PrivateRepo
	GroupRepo   interfaces.GroupRepo
	PublicRepo  interfaces.PublicRepo
}

func NewConferenceServer(authClient auth.AuthClient, monitClient monit.MonitizationClient, pRepo interfaces.PrivateRepo, gRepo interfaces.GroupRepo, puRepo interfaces.PublicRepo) *ConferenceServer {
	return &ConferenceServer{
		AuthClient:  authClient,
		MonitClient: monitClient,
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
	monitReq := &monitPb.Request{
		Data: req.Data,
	}
	resp, err := s.MonitClient.HealthCheck(ctx, monitReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := traceID + resp.Result + req.Data
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
	resp, err := s.AuthClient.GroupHostPermission(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.Permission == false {
		return nil, errors.New("host permission denied")
	}

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
	resp, err := s.AuthClient.PublicHostPermission(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.Permission == false {
		return nil, errors.New("host permission denied")
	}
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
	_ = req.UserID
	response := pb.JoinPrivateConferenceResponse{}
	participantLimit, err := s.PrivateRepo.CheckPrivateLimit(conferenceID)
	if err != nil {

		return nil, err
	}
	currentParticipants, err := s.PrivateRepo.CountPrivateParticipants(conferenceID)
	if err != nil {

		return nil, err
	}
	// permission, err := s.PrivateRepo.CheckPrivateParticipantPermission(conferenceID, userID)
	// if err != nil {

	// 	return nil, err
	// }
	// if permission == false {
	// 	response = pb.JoinPrivateConferenceResponse{
	// 		Result: "Participant permission denied",
	// 	}
	// 	return &response, errors.New("Participant permission denied")
	// }
	if currentParticipants >= participantLimit {
		response = pb.JoinPrivateConferenceResponse{
			Result: "Participant limit exceeded",
		}
		return &response, errors.New("Participant limit exceeded")
	}
	// sdpOffer, err := s.PrivateRepo.GetSdpOffer(conferenceID)
	// if err != nil {
	// 	response = pb.JoinPrivateConferenceResponse{
	// 		Result: "Retrieving sdpoffer from room failed",
	// 	}
	// 	return &response, errors.New("Retrieving sdpoffer from room failed")
	// }
	participantInput := utility.PrivateRoomParticipants{
		Model:        gorm.Model{},
		UserID:       req.UserID,
		ConferenceID: req.ConferenceID,
		SdpAnswer:    "",
		IceCandidate: "",
		Permission:   true,
		CamStatus:    "active",
		MicStatus:    "active",
		JoinTime:     time.Now(),
		ExitTime:     time.Time{},
		Role:         "",
	}
	if err = s.PrivateRepo.AddParticipantInPrivateRoom(participantInput); err != nil {
		response = pb.JoinPrivateConferenceResponse{
			Result: "Adding participant in room failed",
		}
		return &response, errors.New("Adding participant in room failed")
	}
	response = pb.JoinPrivateConferenceResponse{
		Result: "user added to conference room",
	}
	return &response, nil
}

func (s *ConferenceServer) JoinGroupConference(ctx context.Context, req *pb.JoinGroupConferenceRequest) (*pb.JoinGroupConferenceResponse, error) {
	var input utility.GroupRoomParticipants
	traceID := ctx.Value("traceID")
	clientRequest := &authpb.GroupParticipantPermissionRequest{}
	resp, err := s.AuthClient.GroupParticipantPermission(ctx, clientRequest)
	if err != nil {
		return nil, err
	}
	if resp.Permission == false {
		return nil, errors.New("user permission denied")
	}
	participantLimit, err := s.GroupRepo.CheckGroupLimit(req.ConferenceID)
	if err != nil {
		log.Fatal("checklimit", err, ctx.Value("traceID"))
		return nil, err
	}
	currentParticipants, err := s.GroupRepo.CountGroupParticipants(req.ConferenceID)
	if err != nil {
		log.Fatal("countparticipant", err, ctx.Value("traceID"))
		return nil, err
	}
	permission, err := s.GroupRepo.CheckGroupParticipantPermission(req.ConferenceID, req.UserID)
	if err != nil {

		return nil, err
	}
	if permission == false {
		response := pb.JoinGroupConferenceResponse{
			Result: "Participant permission denied",
		}
		return &response, errors.New("Participant permission denied")
	}
	if currentParticipants >= participantLimit {
		response := pb.JoinGroupConferenceResponse{
			Result: "Participant limit exceeded",
		}
		return &response, errors.New("Participant limit exceeded")
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
	clientRequest := &authpb.PublicParticipantPermissionRequest{}
	resp, err := s.AuthClient.PublicParticipantPermission(ctx, clientRequest)
	if err != nil {
		return nil, err
	}
	if resp.Permission == false {
		return nil, errors.New("user permission denied")
	}
	participantLimit, err := s.PublicRepo.CheckPublicLimit(req.ConferenceID)
	if err != nil {
		return nil, err
	}
	currentParticipants, err := s.PublicRepo.CountPublicParticipants(req.ConferenceID)
	if err != nil {
		return nil, err
	}
	permission, err := s.PublicRepo.CheckPublicParticipantPermission(req.ConferenceID, req.UserID)
	if err != nil {

		return nil, err
	}
	if permission == false {
		response := pb.JoinPublicConferenceResponse{
			Result: "Participant permission denied",
		}
		return &response, errors.New("Participant permission denied")
	}
	if currentParticipants >= participantLimit {
		response := pb.JoinPublicConferenceResponse{
			Result: "Participant limit exceeded",
		}
		return &response, errors.New("Participant limit exceeded")
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

func (s *ConferenceServer) RemovePrivateParticipant(ctx context.Context, req *pb.RemovePrivateParticipantRequest) (*pb.RemovePrivateParticipantResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	if req.Block == true {
		if err := s.PrivateRepo.BlockPrivateParticipant(conferenceID, userID); err != nil {
			log.Fatal(err, ctx.Value("traceID"))
			return nil, err
		}
	}
	if err = s.PrivateRepo.RemovePrivateParticipant(conferenceID, userID); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	response := pb.RemovePrivateParticipantResponse{
		Result: "participant removed",
	}
	return &response, nil
}

func (s *ConferenceServer) RemoveGroupParticipant(ctx context.Context, req *pb.RemoveGroupParticipantRequest) (*pb.RemoveGroupParticipantResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	if req.Block == true {
		if err := s.GroupRepo.BlockGroupParticipant(conferenceID, userID); err != nil {
			log.Fatal(err, ctx.Value("traceID"))
			return nil, err
		}
	}
	if err = s.GroupRepo.RemoveGroupParticipant(conferenceID, userID); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	response := pb.RemoveGroupParticipantResponse{
		Result: "participant removed",
	}
	return &response, nil
}

func (s *ConferenceServer) RemovePublicParticipant(ctx context.Context, req *pb.RemovePublicParticipantRequest) (*pb.RemovePublicParticipantResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	if req.Block == true {
		if err := s.PublicRepo.BlockPublicParticipant(conferenceID, userID); err != nil {
			log.Fatal(err, ctx.Value("traceID"))
			return nil, err
		}
	}
	if err = s.PublicRepo.RemovePublicParticipant(conferenceID, userID); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	response := pb.RemovePublicParticipantResponse{
		Result: "participant removed",
	}
	return &response, nil
}

func (s *ConferenceServer) LeavePrivateConference(ctx context.Context, req *pb.LeavePrivateConferenceRequest) (*pb.LeavePrivateConferenceResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	participantInput := utility.PrivateRoomParticipants{
		UserID:       userID,
		ConferenceID: conferenceID,
		ExitTime:     time.Now(),
	}
	if err = s.PrivateRepo.UpdatePrivateParticipantExitTime(participantInput); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	if err = s.PrivateRepo.RemovePrivateParticipant(conferenceID, userID); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	joinTime, err := s.PrivateRepo.GetJoinTime(conferenceID, userID)
	if err != nil {
		log.Println(err, "Get join time err")
	}
	minutes, err := utility.TimeCalculator(joinTime, participantInput.ExitTime)
	if err != nil {
		log.Println(err, "Time calculation failed")
	}
	rewardReq := &monitPb.ParticipationRewardRequest{
		UserID:         userID,
		ConferenceID:   conferenceID,
		ConferenceType: "Private",
		Minutes:        minutes,
	}
	resp, err := s.MonitClient.ParticipationReward(ctx, rewardReq)
	if err != nil {
		log.Println(err, "Monitization server err")
	}
	response := pb.LeavePrivateConferenceResponse{
		Result: "Exited from the conference" + resp,
	}
	return &response, nil
}

func (s *ConferenceServer) LeaveGroupConference(ctx context.Context, req *pb.LeaveGroupConferenceRequest) (*pb.LeaveGroupConferenceResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	participantInput := utility.GroupRoomParticipants{
		UserID:       userID,
		ConferenceID: conferenceID,
		ExitTime:     time.Now(),
	}
	if err = s.GroupRepo.UpdateGroupParticipantExitTime(participantInput); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	if err = s.GroupRepo.RemoveGroupParticipant(conferenceID, userID); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	response := pb.LeaveGroupConferenceResponse{
		Result: "Exited from the conference",
	}
	return &response, nil
}

func (s *ConferenceServer) LeavePublicConference(ctx context.Context, req *pb.LeavePublicConferenceRequest) (*pb.LeavePublicConferenceResponse, error) {
	userID := req.UserID
	conferenceID := req.ConferenceID
	participantInput := utility.PublicRoomParticipants{
		UserID:       userID,
		ConferenceID: conferenceID,
		ExitTime:     time.Now(),
	}
	if err = s.PublicRepo.UpdatePublicParticipantExitTime(participantInput); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	if err = s.PublicRepo.RemovePublicParticipant(conferenceID, userID); err != nil {
		log.Fatal(err, ctx.Value("traceID"))
		return nil, err
	}
	response := pb.LeavePublicConferenceResponse{
		Result: "Exited from the conference",
	}
	return &response, nil
}

func (s *ConferenceServer) EndPrivateConference(ctx context.Context, req *pb.EndPrivateConferenceRequest) (*pb.EndPrivateConferenceResponse, error) {
	err = s.PrivateRepo.RemovePrivateRoom(req.ConferenceID)
	if err != nil {
		return nil, err
	}
	response := pb.EndPrivateConferenceResponse{
		Result: "Conference ended",
	}
	return &response, nil
}

func (s *ConferenceServer) EndGroupConference(ctx context.Context, req *pb.EndGroupConferenceRequest) (*pb.EndGroupConferenceResponse, error) {
	err = s.GroupRepo.RemoveGroupRoom(req.ConferenceID)
	if err != nil {
		return nil, err
	}
	response := pb.EndGroupConferenceResponse{
		Result: "Conference ended",
	}
	return &response, nil
}

func (s *ConferenceServer) EndPublicConference(ctx context.Context, req *pb.EndPublicConferenceRequest) (*pb.EndPublicConferenceResponse, error) {
	err = s.PublicRepo.RemovePublicRoom(req.ConferenceID)
	if err != nil {
		return nil, err
	}
	response := pb.EndPublicConferenceResponse{
		Result: "Conference ended",
	}
	return &response, nil
}

func (s *ConferenceServer) SchedulePrivateConference(ctx context.Context, req *pb.SchedulePrivateConferenceRequest) (*pb.SchedulePrivateConferenceResponse, error) {
	var input utility.ScheduleConference
	ts := &timestamp.Timestamp{
		Seconds: 1694113200,
		Nanos:   0,
	}
	t, err := ptypes.Timestamp(ts)
	if err != nil {
		fmt.Println("Error converting Timestamp:", err)

	}
	copier.Copy(&input, req)
	uid, err := utility.UID(8)
	if err != nil {
		return nil, err
	}
	input.ScheduleID = uid
	input.Time = t
	_, err = s.PrivateRepo.CreatePrivateSchedule(input)
	if err != nil {
		return nil, err
	}
	response := pb.SchedulePrivateConferenceResponse{
		Result:     "Conference scheduled",
		ScheduleID: uid,
	}
	return &response, nil
}

func (s *ConferenceServer) ScheduleGroupConference(ctx context.Context, req *pb.ScheduleGroupConferenceRequest) (*pb.ScheduleGroupConferenceResponse, error) {
	var input utility.ScheduleGroupConference
	if err = copier.Copy(&input, req); err != nil {

	}
	uid, err := utility.UID(8)
	if err != nil {
		return nil, err
	}
	input.ScheduleID = uid
	_, err = s.GroupRepo.CreateGroupSchedule(input)
	if err != nil {
		return nil, err
	}
	response := pb.ScheduleGroupConferenceResponse{
		Result:     "Conference scheduled",
		ScheduleID: uid,
	}
	return &response, nil
}

func (s *ConferenceServer) SchedulePublicConference(ctx context.Context, req *pb.SchedulePublicConferenceRequest) (*pb.SchedulePublicConferenceResponse, error) {
	var input utility.SchedulePublicConference
	copier.Copy(&input, req)
	uid, err := utility.UID(8)
	if err != nil {
		return nil, err
	}
	input.ScheduleID = uid
	_, err = s.PublicRepo.CreatePublicSchedule(input)
	if err != nil {
		return nil, err
	}
	response := pb.SchedulePublicConferenceResponse{
		Result:     "Conference scheduled",
		ScheduleID: uid,
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
