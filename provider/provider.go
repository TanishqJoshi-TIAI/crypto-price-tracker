package provider

import (
	"crypto-price-tracker/config"
	"crypto-price-tracker/responsestruct"
	"fmt"
	"log"
)

const (
	CoinDesk = "CoinDesk"
)

func FetchPrices() (responsestruct.ResponseStruct, error) {
	switch config.APIProvider {
	case CoinDesk:
		return GetPricesFromCoinDeskAPI()
	default:
		log.Println("[WARN] No such provider " + config.APIProvider)
		return responsestruct.ResponseStruct{}, fmt.Errorf("API provider '%s' not supported", config.APIProvider)
	}
}
