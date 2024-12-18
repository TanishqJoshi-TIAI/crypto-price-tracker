package service

import (
	"crypto-price-tracker/config"
	"crypto-price-tracker/provider"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

func CheckAPIExpiry() bool {
	if time.Since(config.LastFetchedAPITime) > config.CacheExpiry {
		return true
	}
	return false
}

func GetCryptoPrice() (gin.H, error) {
	isAPIExpired := CheckAPIExpiry()
	APINotSupported := false
	if isAPIExpired {
		result, err := provider.FetchPrices()
		if err != nil {
			APINotSupported = true
		} else {
			config.CryptoName = result.CryptoName
			config.CryptoPriceInEUR = result.PriceInEUR
			config.CryptoPriceInUSD = result.PriceInUSD
		}
	}

	if APINotSupported {
		return gin.H{}, errors.New("API Provider not supported")
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
