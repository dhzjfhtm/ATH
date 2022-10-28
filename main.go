package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/dhzjfhtm/ATH/realtime/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	binanceClient := api.NewBinanceClient()

	price, err := api.GetBinanceSpotPrice("BTCUSDT", binanceClient.Client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BTC :", price, "$")

	account := binanceClient.GetBinanceAccount()
	balances := account.Balances
	for _, balance := range balances {
		if balance.Free != "0.00000000" {
			fmt.Println(balance.Asset, ":", balance.Free)
		}
	}

	order, err := binanceClient.Client.NewCreateOrderService().Symbol("KLAYUSDT").
		Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity("40").
		Price("0.26").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)

	// Use Test() instead of Do() for testing.
}
