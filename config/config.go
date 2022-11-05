package config

import (
	"github.com/adshao/go-binance/v2/futures"
	"github.com/dhzjfhtm/ATH/realtime/api"
	"github.com/dhzjfhtm/ATH/record"
)

func SetFutureConfig(binanceFuture *api.BinanceFuture, logger *record.Logger) {
	positionRisks, err := binanceFuture.GetAllPositionRisk()
	if err != nil {
		logger.Error("SetFutureConfig", err)
		return
	}

	var commonPositionRisk []*futures.PositionRisk

	for _, symbol := range symbols {
		for _, positionRisk := range positionRisks {
			if symbol == positionRisk.Symbol {
				commonPositionRisk = append(commonPositionRisk, positionRisk)
			}
		}
	}

	for _, positionRisk := range commonPositionRisk {
		if positionRisk.Leverage != "1" {
			binanceFuture.SetLeverage(positionRisk.Symbol, 1)
		}

		if positionRisk.MarginType != "isolated" {
			binanceFuture.SetMarginType(positionRisk.Symbol, "isolated")
		}
	}
}
