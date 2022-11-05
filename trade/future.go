package trade

import (
	"fmt"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/dhzjfhtm/ATH/realtime/api"
)

func FutureTrade() {
	binanceFuture := api.NewBinanceFuture()

	price, err := binanceFuture.GetBinanceFuturePrice("XRPUSDT")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("BTC :", price, "$")

	account := binanceFuture.GetBinanceFutureAccount()
	assets := account.Assets
	for _, asset := range assets {
		if asset.WalletBalance != "0.00000000" {
			fmt.Println(asset.Asset, ":", asset.WalletBalance)
		}
	}

	// check lerverage and margin type
	positionRisk, err := binanceFuture.GetPositionRisk("XRPUSDT")
	if err != nil {
		fmt.Println(err)
		return
	}

	if positionRisk[0].Leverage != "1" {
		binanceFuture.SetLeverage("XRPUSDT", 1)
	}

	if positionRisk[0].MarginType != string(futures.MarginTypeCrossed) {
		binanceFuture.SetMarginType("XRPUSDT", "isolated")
	}

	order, err := binanceFuture.NewBinanceFutureOrder("XRPUSDT", "BUY", "LIMIT", "40", "0.4983")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)

}
