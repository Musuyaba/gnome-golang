package main

import (
	"log"

	"github.com/Musuyaba/gnome-golang/pkg/initializers"
	"github.com/Musuyaba/gnome-golang/routers"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDb(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	routers.ApiRoutes(server)
	routers.PublicRoutes(server)

	log.Fatal(server.RunTLS(config.HOSTNAME+":"+config.PORT, "./certs/generated/server.crt", "./certs/generated/server.key"))
}
