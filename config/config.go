package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
    DatabaseUser     string
    DatabasePassword string
    DatabaseHost     string
    DatabasePort     string
    DatabaseName     string
}

func LoadConfig() Config {
    err := godotenv.Load("config/environment/.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return Config{
        DatabaseUser:     os.Getenv("DATABASE_USER"),
        DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
        DatabaseHost:     os.Getenv("DATABASE_HOST"),
        DatabasePort:     os.Getenv("DATABASE_PORT"),
        DatabaseName:     os.Getenv("DATABASE_NAME"),
    }
}

func GetSecret() string{
    err := godotenv.Load("config/environment/.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    return os.Getenv("SECRET_ACCESSTOKEN")
}

func GetSecretRefresh() string{
    err := godotenv.Load("config/environment/.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    return os.Getenv("SECRET_REFRESHTOKEN")
}

func GetSecretTimeJwt() int64 {
    var defaultTokenTime int64 = 30;
    err := godotenv.Load("config/environment/.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    time,err := strconv.Atoi(os.Getenv("TOKEN_EXPIRED"));
    if err != nil{
        return defaultTokenTime;
    }
    return int64(time);
}