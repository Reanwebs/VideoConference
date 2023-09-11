package monitization

import (
	pb "conference/pb/monitization"
	"conference/pkg/common/config"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type monitizationClient struct {
	Client pb.MonitizationClient
}

func InitClient(c config.Config) (MonitizationClient, error) {

	clientCon, err := grpc.Dial(c.MonitUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println(c.MonitUrl)
	if err != nil {
		return nil, err
	}
	return NewMonitizationClient(pb.NewMonitizationClient(clientCon)), nil
}

func NewMonitizationClient(client pb.MonitizationClient) MonitizationClient {
	return &monitizationClient{
		Client: client,
	}
}

func (a *monitizationClient) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	resp, err := a.Client.HealthCheck(ctx, &pb.Request{
		Data: req.Data,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

func (a *monitizationClient) ParticipationReward(ctx context.Context, req *pb.ParticipationRewardRequest) (string, error) {
	resp, err := a.Client.ParticipationReward(ctx, req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return resp.Result, nil
}
