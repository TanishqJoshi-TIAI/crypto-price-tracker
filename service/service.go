package service

import (
	"crypto-price-tracker/config"
	"crypto-price-tracker/provider"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CheckAPIExpiry() bool {
	if config.LastFetchedAPITime.IsZero() || time.Since(config.LastFetchedAPITime) > config.CacheExpiry {
		return true
	}
	return false
}

func GetCryptoPrice() (gin.H, error) {
	isAPIExpired := CheckAPIExpiry()
	if isAPIExpired {
		result, err := provider.FetchPrices()
		if err != nil {
			return gin.H{}, fmt.Errorf("failed to fetch cryptocurrency prices: %w", err)
		} else {
			config.CryptoName = result.CryptoName
			config.CryptoPriceInEUR = result.PriceInEUR
			config.CryptoPriceInUSD = result.PriceInUSD
		}
	}

	return gin.H{
		"data": gin.H{
			config.CryptoName: gin.H{
				"EUR": config.CryptoPriceInEUR,
				"USD": config.CryptoPriceInUSD,
			},
		},
	}, nil
}
