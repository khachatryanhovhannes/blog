package main

import (
	"backend/database"
	"backend/models" // Import models
	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()

	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Post{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow your React app origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Optional: allow credentials (cookies, authorization headers, etc.)
	}))

	routes.RegisterRoutes(r)

	r.Run()
}
