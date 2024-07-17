# go_project_structure

# PACKAGE INSTALLED
go mod init go_project_structure
go get github.com/joho/godotenv
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

# LIBRARY
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt

# SETUP FOLDER STEP

Config
    - config.go
    environment
        - .env

database
    - database.go

routes
    - routes.go (main route)

internal 
    - handler
    - models
    - repository
    - routes
    - services

middleware
    - auth_middleware.go for intercept request
 
utils