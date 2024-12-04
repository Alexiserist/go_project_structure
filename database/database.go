package database

import (
	"fmt"
	"go_project_structure/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error{
	config := config.LoadConfig();
	log.Print("Connecting To Database:", config.DatabaseHost, ":", config.DatabasePort);
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, config.DatabaseName);

	db,err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to Database: %v", err)
		return err
	}
	DB = db;

	sqlDB, err := DB.DB();
	if err != nil {
		return err
	} 
	sqlDB.SetMaxIdleConns(10);
	sqlDB.SetMaxOpenConns(100);
	sqlDB.SetConnMaxLifetime(time.Minute * 5);

	return nil
}	
