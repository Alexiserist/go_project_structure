package main

import (
	"log"
	
	"go_project_structure/database"
	"go_project_structure/internal/routes"
	_ "go_project_structure/docs"
)

// @title           Golang Framework GIN Swagger
// @version         1.0
// @description     Golang Framwork GIN ,Database Postgress, Gorm, Swaggo

// @host      localhost:8080

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main()  {
	err := database.Init();
	if err != nil {
	   log.Fatalf("Could not initialize the database: %v", err)
	   return
   }
   
   router := routes.LoadRouter();
   if err := router.Run(":8080"); err != nil {
	   log.Fatalf("Could not start the server: %v", err)
   }
}