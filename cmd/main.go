package main

import (
	"crypto-price-tracker/config"
	"crypto-price-tracker/middleware"
	"crypto-price-tracker/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	route := gin.Default()
	route.Use(middleware.Logger())
	route.GET("/prices", func(c *gin.Context) {
		response, err := service.GetCryptoPrice()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		} else {
			c.JSON(http.StatusOK, response)
		}
	})
	err := route.Run(config.LocalHostPort)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}
