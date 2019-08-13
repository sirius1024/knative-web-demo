package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/dog", func(c *gin.Context) {
		c.String(http.StatusOK, "DOG!")
	})

	r.POST("/cat", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "miao u~~",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
