package api

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2"
)

var (
	apiKey    = os.Getenv("BINANCE_API_KEY")
	secretKey = os.Getenv("BINANCE_API_SECRET")
)

type BinanceClient struct {
	Client *binance.Client
}

func NewBinanceClient() *BinanceClient {
	if apiKey == "" || secretKey == "" {
		apiKey = os.Getenv("BINANCE_API_KEY")
		secretKey = os.Getenv("BINANCE_API_SECRET")
	}

	return &BinanceClient{
		Client: binance.NewClient(apiKey, secretKey),
	}
}

func GetBinanceSpotPrice(coin string, client *binance.Client) (string, error) {
	// get binance spot price
	tickerPrice, err := client.NewListPricesService().Symbol(coin).Do(context.Background())
	if err != nil {
		return "", err
	}
	return tickerPrice[0].Price, nil
}
