package main

import (
	"fmt"
	"os"

	"github.com/AlejandroWaiz/http-archetype/controller"
	"github.com/AlejandroWaiz/http-archetype/resttemplate"
	"github.com/AlejandroWaiz/http-archetype/router"
	"github.com/AlejandroWaiz/http-archetype/server"
	"github.com/AlejandroWaiz/http-archetype/service"
	"github.com/joho/godotenv"
)

func main() {

	var Version = "1.0"
	fmt.Printf("API Http Adapter App Version: %s\n", Version)
	//Loads environment variables from .env file if exists
	godotenv.Load()

	rest := resttemplate.GetIRestTemplate()
	service := service.GetIService(rest)
	controller := controller.GetIController(service)
	router := router.GetIRouter(controller)

	server := server.New()

	router.RegisterRoutes(server)

	server.Run(":" + os.Getenv("HTTP_ADAPTER_PORT"))

}
