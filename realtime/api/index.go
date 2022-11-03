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

type BinanceClient struct {
	client *binance.Client
}

func NewBinanceClient() *BinanceClient {
	if apiKey == "" || secretKey == "" {
		apiKey = os.Getenv("BINANCE_API_KEY")
		secretKey = os.Getenv("BINANCE_API_SECRET")
	}

	return &BinanceClient{
		client: binance.NewClient(apiKey, secretKey),
	}
}

// Get Account
func (bc *BinanceClient) GetBinanceAccount() *binance.Account {
	balances, err := bc.client.NewGetAccountService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	return balances
}

func (bc *BinanceClient) GetBinanceSpotPrice(coin string) (string, error) {
	// get binance spot price
	tickerPrice, err := bc.client.NewListPricesService().Symbol(coin).Do(context.Background())
	if err != nil {
		return "", err
	}
	return tickerPrice[0].Price, nil
}

func (bc *BinanceClient) NewBinanceSpotOrder(symbol, side, orderType, quantity, price string) (*binance.CreateOrderResponse, error) {
	order, err := bc.client.NewCreateOrderService().Symbol(symbol).
		Side(binance.SideType(side)).Type(binance.OrderType(orderType)).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity(quantity).
		Price(price).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}
