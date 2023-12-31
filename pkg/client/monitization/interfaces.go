package monitization

import (
	pb "conference/pb/monitization"
	"context"
)

type MonitizationClient interface {
	HealthCheck(context.Context, *pb.Request) (*pb.Response, error)
	ParticipationReward(context.Context, *pb.ParticipationRewardRequest) (string, error)
}
