package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HOLA GIN",
		})
	})

	r.POST("/users/", func(c *gin.Context) {
		name := c.PostForm("name")
		lastname := c.PostForm("lastname")
		c.JSON(200, gin.H{ // serializador de gin
			"name":     name,
			"lastname": lastname,
		})
	})

	r.Run()
}
