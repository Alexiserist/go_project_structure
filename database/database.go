package database

import (
	"fmt"
	"go_project_structure/config"
	"log"

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
	// MigrateDB(db);
	DB = db;
	return nil
}	

// func MigrateDB(db *gorm.DB){
// 	db.AutoMigrate(&models.User{});
// }