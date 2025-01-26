package main

import (
	"log"

	"github.com/gin-gonic/gin"
	routes "github.com/giselescarvalho/CRUDemGo/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	// Chamada correta para a função InitRoutes
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
