package api

import (
	"context"

	"github.com/adshao/go-binance/v2"
)

func GetBinanceSpotPrice(coin string) (string, error) {
	// get binance spot price
	client := binance.NewClient("", "")
	tickerPrice, err := client.NewListPricesService().Symbol(coin).Do(context.Background())
	if err != nil {
		return "", err
	}
	return tickerPrice[0].Price, nil
}
