package database

import (
	"conference/pkg/common/config"
	"conference/pkg/common/utility"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(cfg config.Config) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DbHost, cfg.DbUser, cfg.DbName, cfg.DbPort, cfg.DbPassword)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to database:", err)
		return nil
	}
	DB = db
	DB.AutoMigrate(utility.ConferenceRoom{})
	return db
}
