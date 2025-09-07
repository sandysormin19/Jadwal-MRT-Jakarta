package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sandysormin19/Jadwal-MRT/modules/station"
	"os"
	"time"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()

	// Konfigurasi CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // development
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes
	api := router.Group("/v1/api")
	station.Initiate(api)

	// Serve frontend React build
	router.Static("/static", "./frontend-build/static")
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend-build/index.html")
	})

	// Port dinamis
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
