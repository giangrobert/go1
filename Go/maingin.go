package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HOLA GIN",
		})
	})

	// For each matched request Context will hold the route definition
	r.POST("/users/", func(c *gin.Context) {
		id := c.Query("id")
		name := c.PostForm("name")
		c.JSON(200, gin.H{
			"message": name + " id=" + id,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
