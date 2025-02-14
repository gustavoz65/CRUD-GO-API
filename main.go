package main

import (
	"log"

	"github.com/gin-gonic/gin"
	routes "github.com/gustavoz65/CRUD-GO-API.git/src/controller/routesUser"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
}
