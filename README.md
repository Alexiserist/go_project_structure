# GO Project Structure

# Package Installed

```bash
  go mod init go_project_structure
```

```bash
  go get github.com/joho/godotenv
```

```bash
  go get -u github.com/gin-gonic/gin
```

```bash
  go get -u gorm.io/gorm
```

```bash
  go get -u gorm.io/driver/postgres
```

```bash
  go get -u github.com/swaggo/swag/cmd/swag
```

```bash
  go get -u github.com/swaggo/gin-swagger
```

```bash
  go get -u github.com/swaggo/files
```

# Library

```bash
  go get github.com/golang-jwt/jwt/v5
```

```bash
  go get golang.org/x/crypto/bcrypt
```

# Setup Folder Setup

Config
- config.go
- environment
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
- auth_middleware.go (For intercept request)
 
utils