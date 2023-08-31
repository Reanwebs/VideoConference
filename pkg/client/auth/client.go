package auth

import (
	pb "conference/pb/auth"
	"conference/pkg/common/config"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	Client pb.AuthClient
}

func InitClient(c config.Config) (AuthClient, error) {
	clientCon, err := grpc.Dial(c.AuthUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewAuthClient(pb.NewAuthClient(clientCon)), nil
}

func NewAuthClient(client pb.AuthClient) AuthClient {
	return &authClient{
		Client: client,
	}
}

func (a *authClient) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	resp, err := a.Client.HealthCheck(ctx, &pb.Request{
		Data: "Hi authentication server",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) GroupHostPermission(ctx context.Context, req *pb.GroupHostPermissionRequest) (*pb.GroupHostPermissionResponse, error) {
	resp := &pb.GroupHostPermissionResponse{
		Permission: true,
	}

	return resp, nil
}

func (a *authClient) GroupParticipantPermission(ctx context.Context, req *pb.GroupParticipantPermissionRequest) (*pb.GroupParticipantPermissionResponse, error) {
	resp := &pb.GroupParticipantPermissionResponse{
		Permission: true,
	}

	return resp, nil
}

func (a *authClient) PublicHostPermission(ctx context.Context, req *pb.PublicHostPermissionRequest) (*pb.PublicHostPermissionResponse, error) {
	resp := &pb.PublicHostPermissionResponse{
		Permission: true,
	}

	return resp, nil
}

func (a *authClient) PublicParticipantPermission(ctx context.Context, req *pb.PublicParticipantPermissionRequest) (*pb.PublicParticipantPermissionResponse, error) {
	resp := &pb.PublicParticipantPermissionResponse{
		Permission: true,
	}

	return resp, nil
}
