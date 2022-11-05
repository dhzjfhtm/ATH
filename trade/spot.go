package trade

import (
	"fmt"

	"github.com/dhzjfhtm/ATH/realtime/api"
)

func SpotTrade() {
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
}
