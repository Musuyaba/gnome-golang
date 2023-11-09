package main

import (
	"fmt"
	"log"

	"github.com/Musuyaba/gnome-golang/pkg/initializers"
	"github.com/Musuyaba/gnome-golang/pkg/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	initializers.ConnectDbGormPostgreSQL(&config)
}

func main() {
	initializers.PostgreGormInstance.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
