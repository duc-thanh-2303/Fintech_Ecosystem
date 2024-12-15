package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Tạo router
	r := gin.Default()

	r.GET("/", func(c *ginContext) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World"
		})
	}

}
