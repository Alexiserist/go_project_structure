package routes

import (
	"go_project_structure/database"
	"go_project_structure/internal/handler"
	"go_project_structure/internal/repository"
	"go_project_structure/internal/services"
	"go_project_structure/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	authMiddleware := middleware.NewAuthMiddleware()

	userRepository := repository.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	users := router.Group("/users")
	users.Use(authMiddleware.AuthorizationMiddleware)
	users.GET("/", userHandler.GetUsers)
	users.GET("/:id", userHandler.GetUserByKey)
	users.POST("/", userHandler.CreateUser)
	users.DELETE("/:id", userHandler.DeleteUser)
}
