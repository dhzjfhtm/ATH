package api

import (
	"context"
	"fmt"
	"os"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/dhzjfhtm/ATH/record"
)

type BinanceFuture struct {
	client  *futures.Client
	verbose bool
	logger  *record.Logger
}

func NewBinanceFuture(logger *record.Logger) *BinanceFuture {
	if apiKey == "" || secretKey == "" {
		apiKey = os.Getenv("BINANCE_API_KEY")
		secretKey = os.Getenv("BINANCE_API_SECRET")
	}

	return &BinanceFuture{
		client:  futures.NewClient(apiKey, secretKey),
		verbose: true,
		logger:  logger,
	}
}

// get binance future account
func (bf *BinanceFuture) GetBinanceFutureAccount() *futures.Account {
	balances, err := bf.client.NewGetAccountService().Do(context.Background())
	bf.logger.Info("GetBinanceFutureAccount", fmt.Sprintf("%+v", balances))
	if err != nil {
		bf.logger.Error("GetBinanceFutureAccount", err)
		panic(err)
	}

	return balances
}

func (bf *BinanceFuture) GetBinanceFuturePrice(symbol string) (string, error) {
	// get binance spot price
	tickerPrice, err := bf.client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		bf.logger.Error("GetBinanceFuturePrice", err)
		return "", err
	}
	return tickerPrice[0].Price, nil
}

// get binance klines data
func (bf *BinanceFuture) GetBinanceFutureKlines(symbol, interval string, limit int) ([]*futures.Kline, error) {
	klines, err := bf.client.NewKlinesService().Symbol(symbol).Interval(interval).Limit(limit).Do(context.Background())
	if err != nil {
		bf.logger.Error("GetBinanceFutureKlines", err)
		return nil, err
	}

	return klines, nil
}

func (bf *BinanceFuture) NewBinanceFutureOrder(symbol, side, orderType, quantity, price string) (*futures.CreateOrderResponse, error) {
	order, err := bf.client.NewCreateOrderService().Symbol(symbol).
		Side(futures.SideType(side)).Type(futures.OrderType(orderType)).
		TimeInForce(futures.TimeInForceTypeGTC).Quantity(quantity).
		Price(price).Do(context.Background())
	bf.logger.Info("NewBinanceFutureOrder", fmt.Sprintf("%+v", order))
	if err != nil {
		bf.logger.Error("NewBinanceFutureOrder", err)
		return nil, err
	}

	return order, nil
}

// set leverage
func (bf *BinanceFuture) SetLeverage(symbol string, leverage int) error {
	_, err := bf.client.NewChangeLeverageService().Symbol(symbol).Leverage(leverage).Do(context.Background())
	bf.logger.Info("SetLeverage", fmt.Sprintf("%+v", leverage))
	if err != nil {
		bf.logger.Error("SetLeverage", err)
		return err
	}

	return nil
}

// set margin type
func (bf *BinanceFuture) SetMarginType(symbol string, marginType futures.MarginType) error {
	err := bf.client.NewChangeMarginTypeService().Symbol(symbol).MarginType(marginType).Do(context.Background())
	bf.logger.Info("SetMarginType", fmt.Sprintf("%+v", marginType))
	if err != nil {
		bf.logger.Error("SetMarginType", err)
		return err
	}
	return nil
}

// get position risk
func (bf *BinanceFuture) GetPositionRisk(symbol string) ([]*futures.PositionRisk, error) {
	positionRisk, err := bf.client.NewGetPositionRiskService().Symbol(symbol).Do(context.Background())
	if err != nil {
		bf.logger.Error("GetPositionRisk", err)
		return nil, err
	}

	return positionRisk, nil
}

// get all position risks
func (bf *BinanceFuture) GetAllPositionRisk() ([]*futures.PositionRisk, error) {
	positionRisk, err := bf.client.NewGetPositionRiskService().Do(context.Background())
	if err != nil {
		bf.logger.Error("GetAllPositionRisk", err)
		return nil, err
	}

	return positionRisk, nil
}
