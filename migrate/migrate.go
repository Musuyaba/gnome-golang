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
	initializers.ConnectDbPostgreSQL(&config)
}

func main() {
	initializers.DbInstance.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
