package api

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2/futures"
)

type BinanceFuture struct {
	client *futures.Client
}

func NewBinanceFuture() *BinanceFuture {
	if apiKey == "" || secretKey == "" {
		apiKey = os.Getenv("BINANCE_API_KEY")
		secretKey = os.Getenv("BINANCE_API_SECRET")
	}

	return &BinanceFuture{
		client: futures.NewClient(apiKey, secretKey),
	}
}

// get binance future account
func (bf *BinanceFuture) GetBinanceFutureAccount() *futures.Account {
	balances, err := bf.client.NewGetAccountService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	return balances
}

func (bf *BinanceFuture) GetBinanceFuturePrice(symbol string) (string, error) {
	// get binance spot price
	tickerPrice, err := bf.client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return "", err
	}
	return tickerPrice[0].Price, nil
}

func (bf *BinanceFuture) NewBinanceFutureOrder(symbol, side, orderType, quantity, price string) (*futures.CreateOrderResponse, error) {
	order, err := bf.client.NewCreateOrderService().Symbol(symbol).
		Side(futures.SideType(side)).Type(futures.OrderType(orderType)).
		TimeInForce(futures.TimeInForceTypeGTC).Quantity(quantity).
		Price(price).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}

// set leverage
func (bf *BinanceFuture) SetLeverage(symbol string, leverage int) error {
	_, err := bf.client.NewChangeLeverageService().Symbol(symbol).Leverage(leverage).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// set margin type
func (bf *BinanceFuture) SetMarginType(symbol string, marginType futures.MarginType) error {
	err := bf.client.NewChangeMarginTypeService().Symbol(symbol).MarginType(marginType).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
