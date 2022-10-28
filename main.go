package main

import (
	"fmt"

	"github.com/dhzjfhtm/ATH/realtime/api"
)

func main() {
	fmt.Println("BTC : 20000$")
	price, err := api.GetBinanceSpotPrice("BTCUSDT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BTC :", price, "$")
}
