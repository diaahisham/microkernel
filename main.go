package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"microkernel/core"
)

func main() {
	log.Println("starting the server")
	server := gin.Default()
	server.POST("/pay", CORSMiddleware(), core.Pay)
	server.POST("/refund", CORSMiddleware(), core.Refund)
	_ = server.Run(":7070")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Correlation-Id")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")

	}
}
