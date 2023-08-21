package di

import (
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

	repo := repo.NewConferenceRepo(DB)

	usecase := service.NewConferenceServer(repo)

	server := service.NewGrpcServer(cfg, usecase)

	return server, nil

}
