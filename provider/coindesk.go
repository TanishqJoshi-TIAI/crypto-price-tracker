package provider

import (
	"crypto-price-tracker/config"
	"crypto-price-tracker/responsestruct"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type CoinDeskResponse struct {
	Time       TimeInfo `json:"time"`
	Disclaimer string   `json:"disclaimer"`
	ChartName  string   `json:"chartName"`
	Bpi        BpiInfo  `json:"bpi"`
}

type TimeInfo struct {
	Updated    string `json:"updated"`
	UpdatedISO string `json:"updatedISO"`
	UpdatedUK  string `json:"updateduk"`
}

type BpiInfo struct {
	USD CurrencyInfo `json:"USD"`
	EUR CurrencyInfo `json:"EUR"`
	GBP CurrencyInfo `json:"GBP"`
}

type CurrencyInfo struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}

func GetPricesFromCoinDeskAPI() responsestruct.ResponseStruct {
	res, err := http.Get(config.CoinDeskAPIProviderURL)
	config.LastFetchedAPITime = time.Now()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response CoinDeskResponse
	var result responsestruct.ResponseStruct
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	result.CryptoName = response.ChartName
	result.PriceInEUR = response.Bpi.EUR.Rate
	result.PriceInUSD = response.Bpi.USD.Rate
	return result
}
