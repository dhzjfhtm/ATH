package api

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2"
)

var (
	apiKey    = os.Getenv("BINANCE_API_KEY")
	secretKey = os.Getenv("BINANCE_API_SECRET")
)

type BinanceSpot struct {
	client *binance.Client
}

func NewBinanceSpot() *BinanceSpot {
	if apiKey == "" || secretKey == "" {
		apiKey = os.Getenv("BINANCE_API_KEY")
		secretKey = os.Getenv("BINANCE_API_SECRET")
	}

	return &BinanceSpot{
		client: binance.NewClient(apiKey, secretKey),
	}
}

// Get Account
func (bs *BinanceSpot) GetBinanceSpotAccount() *binance.Account {
	balances, err := bs.client.NewGetAccountService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	return balances
}

func (bs *BinanceSpot) GetBinanceSpotPrice(coin string) (string, error) {
	// get binance spot price
	tickerPrice, err := bs.client.NewListPricesService().Symbol(coin).Do(context.Background())
	if err != nil {
		return "", err
	}
	return tickerPrice[0].Price, nil
}

func (bs *BinanceSpot) NewBinanceSpotOrder(symbol, side, orderType, quantity, price string) (*binance.CreateOrderResponse, error) {
	order, err := bs.client.NewCreateOrderService().Symbol(symbol).
		Side(binance.SideType(side)).Type(binance.OrderType(orderType)).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity(quantity).
		Price(price).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}
