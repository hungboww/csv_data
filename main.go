package main

import (
	"ads/CSV"
	"ads/database"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDatabase()
	CSV.ReadCSV()
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
