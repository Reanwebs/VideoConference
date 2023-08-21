package main

import (
	"conference/pkg/common/config"
	"conference/pkg/common/di"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatalln(err)
	} else {
		if err := server.StartServer(config); err != nil {
			log.Fatalln(err)
		}

	}

}
