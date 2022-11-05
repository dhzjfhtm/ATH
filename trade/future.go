package trade

import (
	"fmt"

	"github.com/dhzjfhtm/ATH/realtime/api"
)

func FutureTrade() {
	binanceFuture := api.NewBinanceFuture()

	price, err := binanceFuture.GetBinanceFuturePrice("BTCUSDT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BTC :", price, "$")

	account := binanceFuture.GetBinanceFutureAccount()
	assets := account.Assets
	for _, asset := range assets {
		if asset.WalletBalance != "0.00000000" {
			fmt.Println(asset.Asset, ":", asset.WalletBalance)
		}
	}
}
