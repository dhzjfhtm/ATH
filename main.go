package main

import (
	"fmt"

	"github.com/dhzjfhtm/ATH/realtime/api"
)

func main() {
	binanceClient := api.NewBinanceClient()

	price, err := api.GetBinanceSpotPrice("BTCUSDT", binanceClient)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BTC :", price, "$")
}
