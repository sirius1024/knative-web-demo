package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ROOT!!! v2")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong(v2)",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "(v2)Hello %s", name)
	})

	r.GET("/dog", func(c *gin.Context) {
		c.String(http.StatusOK, "DOG!(v2)")
	})

	r.POST("/cat", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "miao u~~(v2)",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
