package provider

import (
	"crypto-price-tracker/config"
	"crypto-price-tracker/responsestruct"
	"errors"
)

const (
	CoinDesk = "CoinDesk"
)

func FetchPrices() (responsestruct.ResponseStruct, error) {
	switch config.APIProvider {
	case CoinDesk:
		return GetPricesFromCoinDeskAPI(), nil
	default:
		return responsestruct.ResponseStruct{}, errors.New("API provider not supported")
	}
}
