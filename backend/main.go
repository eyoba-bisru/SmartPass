package main

import (
	"github.com/eyoba-bisru/SmartPass/db"
	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectDB()
	defer db.CloseDB()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
