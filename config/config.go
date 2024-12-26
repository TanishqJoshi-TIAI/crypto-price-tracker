package config

import "time"

const (
	CacheExpiry            = 1 * time.Minute
	CoinDeskAPIProviderURL = "https://api.coindesk.com/v1/bpi/currentprice.json"
	LocalHostPort          = ":8080"
)

var (
	LastFetchedAPITime time.Time
	APIProvider        = "CoinDesk"
	CryptoName         string
	CryptoPriceInEUR   string
	CryptoPriceInUSD   string
)
