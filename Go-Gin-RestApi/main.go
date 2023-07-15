package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/submit", func(c *gin.Context) {
		var person Person

		if err := c.ShouldBind(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name": person.Name,
		})
	})

	router.Run()
}
