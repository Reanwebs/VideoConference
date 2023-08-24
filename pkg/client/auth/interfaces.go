package auth

import (
	pb "conference/pb/auth"
	"context"
)

type AuthClient interface {
	HealthCheck(context.Context, *pb.Request) (*pb.Response, error)
	GroupHostPermission(ctx context.Context, req *pb.GroupHostPermissionRequest) (*pb.GroupHostPermissionResponse, error)
	GroupParticipantPermission(context.Context, *pb.GroupParticipantPermissionRequest) (*pb.GroupParticipantPermissionResponse, error)
	PublicHostPermission(context.Context, *pb.PublicHostPermissionRequest) (*pb.PublicHostPermissionResponse, error)
	PublicParticipantPermission(context.Context, *pb.PublicParticipantPermissionRequest) (*pb.PublicParticipantPermissionResponse, error)
}
