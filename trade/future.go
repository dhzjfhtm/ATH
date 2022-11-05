package trade

import (
	"fmt"

	"github.com/dhzjfhtm/ATH/realtime/api"
)

func FutureTrade(binanceFuture *api.BinanceFuture) {
	price, err := binanceFuture.GetBinanceFuturePrice("XRPUSDT")
	if err != nil {
		fmt.Println(err)
		return
	}

	account := binanceFuture.GetBinanceFutureAccount()
	assets := account.Assets
	for _, asset := range assets {
		if asset.WalletBalance != "0.00000000" {
			fmt.Println(asset.Asset, ":", asset.WalletBalance)
		}
	}

	order, err := binanceFuture.NewBinanceFutureOrder("XRPUSDT", "BUY", "LIMIT", "40", price)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
}
