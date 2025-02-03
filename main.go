package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/giselescarvalho/CRUDemGo/src/configuration/logger"
	"github.com/giselescarvalho/CRUDemGo/src/controller"
	"github.com/giselescarvalho/CRUDemGo/src/controller/model/service"
	routes "github.com/giselescarvalho/CRUDemGo/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("##### ---- ##### About to start application ##### ---- ##### ")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init dependecies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	// Chamada correta para a função InitRoutes
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
