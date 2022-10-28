package main

import (
	"fmt"

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
}
