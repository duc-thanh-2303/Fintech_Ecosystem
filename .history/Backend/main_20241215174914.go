package main

import (
	"net/http"

	"backend/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	//Tạo router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Kết nối database
	configs.ConnectDatabase()

	//Khởi chạy server
	r.Run(":8000")

}
