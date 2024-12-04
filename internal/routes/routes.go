package routes

import (
	"go_project_structure/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadRouter() *gin.Engine {
	router := gin.Default();
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTION"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	config.AllowCredentials = true;
	config.AllowOrigins = []string{"http://localhost:4200"}
	router.Use(cors.New(config))
	auth.AuthRoute(router)
	UserRoute(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
