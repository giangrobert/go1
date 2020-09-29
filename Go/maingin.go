package main

import (
	//"model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	dsn := "docker:docker@tcp(db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Person{})

	//s := Person{Name: "Sean", Age: 50}
	//s.Name = "Sean"

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
