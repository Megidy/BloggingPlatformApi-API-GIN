package main

import (
	"github.com/Megidy/BloggingPlatform-Api/pkj/config"
	"github.com/Megidy/BloggingPlatform-Api/pkj/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	r := gin.Default()
	r.POST("/posts/", controllers.CreatePost)
	r.GET("/posts/", controllers.GetPosts)
	r.GET("/posts/:postId", controllers.GetPostById)
	r.DELETE("/posts/:postId", controllers.DeletePost)
	r.PUT("/posts/:postId", controllers.UpdatePost)

	r.Run(":8080")
}
