package di

import (
	auth "conference/pkg/client/auth"
	monit "conference/pkg/client/monitization"
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

	authClient, err := auth.InitClient(cfg)
	if err != nil {
		return nil, err
	}

	monitClient, err := monit.InitClient(cfg)
	if err != nil {
		return nil, err
	}

	usecase := service.NewConferenceServer(authClient, monitClient, PrivateRepo, GroupRepo, PublicRepo)

	server := service.NewGrpcServer(cfg, usecase)

	return server, nil

}
