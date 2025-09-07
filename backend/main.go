package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sandysormin19/Jadwal-MRT/modules/station"
	"time"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()

	// Konfigurasi CORS untuk API saja
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // untuk production bisa diganti domain tertentu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes
	api := router.Group("/v1/api")
	station.Initiate(api)

	// Serve frontend React build
	router.Static("/static", "./frontend-build/static") // folder static
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend-build/index.html") // selalu fallback ke index.html
	})

	// Jalankan server
	router.Run(":8080")
}
