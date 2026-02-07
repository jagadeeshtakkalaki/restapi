package controllers

import "github.com/gin-gonic/gin"
import "log"


func Ping(c *gin.Context) {
	log.Printf("Ping handler executed")
    c.JSON(200, gin.H{"message": "pong"})
}
