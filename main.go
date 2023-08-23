package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"log"
	"soundcase/controllers"
)

func init() {
	chi.RegisterMethod("PUT")
	chi.RegisterMethod("POST")
}

var (
	ListenAddr = "localhost:8081"
	server     *gin.Engine
)

func init() {
	log.Printf("Starting up on %s", ListenAddr)

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	fmt.Println("Starting application ...")
	fmt.Printf("Config application %v ...", config)

	r := gin.Default()
	r.GET("/api/artist/:id", controllers.ReadArtist)
	r.GET("/api/artist", controllers.ReadArtists)
	r.POST("/api/artist", controllers.CreateArtist)
	r.PUT("/api/artist/:id", controllers.UpdateArtist)
	r.DELETE("/api/artist/:id", controllers.DeleteArtist)

	r.GET("/api/track/:id", controllers.ReadTrack)
	r.GET("/api/track", controllers.ReadTracks)
	r.POST("/api/track", controllers.CreateTrack)
	r.PUT("/api/track/:id", controllers.UpdateTrack)
	r.DELETE("/api/track/:id", controllers.DeleteTrack)

	r.GET("/api/files", controllers.ReadFile)

	err = r.Run(":8081")
	if err != nil {
		fmt.Printf("There are some problems when trying to run app...")
		return
	}
}
