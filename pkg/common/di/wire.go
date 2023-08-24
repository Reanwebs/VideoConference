package di

import (
	client "conference/pkg/client/auth"
	"conference/pkg/common/config"
	db "conference/pkg/repository/database"
	"conference/pkg/repository/repo"
	"conference/pkg/service"
)

// ff go:build wireinject
// ff +build wireinject
// func InitializeAPI(cfg config.Config) (*service.Server, error) {
// 	wire.Build(
// 		db.ConnectToDB,
// 		repo.NewConferenceRepo,
// 		service.NewConferenceServer,
// 		service.NewGrpcServer,
// 	)
// 	return nil, nil
// }

func InitializeAPI(cfg config.Config) (*service.Server, error) {
	DB := db.ConnectToDB(cfg)

	PrivateRepo := repo.NewPrivateConferenceRepo(DB)
	GroupRepo := repo.NewGroupConferenceRepo(DB)
	PublicRepo := repo.NewPublicConferenceRepo(DB)

	client, err := client.InitClient(cfg)
	if err != nil {
		return nil, err
	}

	usecase := service.NewConferenceServer(client, PrivateRepo, GroupRepo, PublicRepo)

	server := service.NewGrpcServer(cfg, usecase)

	return server, nil

}
