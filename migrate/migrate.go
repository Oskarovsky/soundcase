package main

import (
	"fmt"
	"log"
	models "soundcase/models"

	"github.com/wpcodevo/golang-gorm-postgres/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(models.Track{}, models.Artist{})
	fmt.Println("? Migration complete")
}
