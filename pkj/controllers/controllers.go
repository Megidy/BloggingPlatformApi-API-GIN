package controllers

import (
	"net/http"
	"strconv"

	"github.com/Megidy/BloggingPlatform-Api/pkj/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	//retrieving data form database
	posts, err := models.GetAllPosts()
	if err != nil {
		//hadnling error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "didnt retrieve data from database",
		})
	}
	//response to client
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
func CreatePost(c *gin.Context) {
	var NewPost models.Post
	err := c.ShouldBindJSON(&NewPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
	}
	createdPost, err := models.CreatePost(&NewPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "database didnt recieve data",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"post:": createdPost,
	})

}
func GetPostById(c *gin.Context) {

	id := c.Param("postId")
	postId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt convert post id to int",
		})
	}
	post, _, err := models.GetPostById(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt retrieve data from database",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("postId")
	postId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt convert post id to int",
		})
	}
	post, err := models.DeletePost(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt retrieve ro delete data",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"deleted post": post,
	})
}
func UpdatePost(c *gin.Context) {
	var NewPost models.Post

	err := c.ShouldBindJSON(&NewPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
	}

	id := c.Param("postId")
	postId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt convert post id to int",
		})
	}
	_, err = models.DeletePost(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt retrieve ro delete data",
		})
	}
	post, err := models.CreatePost(&NewPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Updated Post:": post,
	})
}
