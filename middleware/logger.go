package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.Printf("Incoming Request: Method = %s, Request URI = %s", c.Request.Method, c.Request.RequestURI)
		c.Next()
		latency := time.Since(start)
		log.Printf("Outgoing Response: Latency = %s, Client IP = %s, Status Code = %d", latency, c.ClientIP(), c.Writer.Status())
	}
}
