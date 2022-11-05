package main

import (
	"github.com/dhzjfhtm/ATH/record"
)

func main() {
	/*
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}

		binanceSpot := api.NewBinanceSpot()

		price, err := binanceSpot.GetBinanceSpotPrice("BTCUSDT")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("BTC :", price, "$")

		account := binanceSpot.GetBinanceSpotAccount()
		balances := account.Balances
		for _, balance := range balances {
			if balance.Free != "0.00000000" {
				fmt.Println(balance.Asset, ":", balance.Free)
			}
		}

		order, err := binanceSpot.NewBinanceSpotOrder("KLAYUSDT", "BUY", "LIMIT", "40", "0.26")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(order)

		// Use Test() instead of Do() for testing.
	*/
	logger := record.NewLogger()

	logger.Info("Hello World")
}
