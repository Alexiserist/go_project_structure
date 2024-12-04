package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine){
	authRepository := NewAuthRepository();
	authService := NewAuthService(authRepository);
	authHandler := NewAuthHandler(authService);
	router.POST("/auth/login", authHandler.Login);
	router.POST("/auth/refresh", authHandler.RefreshToken);


}