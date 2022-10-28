package api

import (
	"context"

	"github.com/adshao/go-binance/v2"
)

type BinanceClient struct {
	Client *binance.Client
}

func NewBinanceClient() *BinanceClient {
	return &BinanceClient{
		Client: binance.NewClient("", ""),
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
