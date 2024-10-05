package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())

	authorized.POST("/posts", controllers.CreatePost)
	authorized.PUT("/posts/:id", controllers.UpdatePost)
	authorized.DELETE("/posts/:id", controllers.DeletePost)
}
