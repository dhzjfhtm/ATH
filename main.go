package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/dhzjfhtm/ATH/config"
	"github.com/dhzjfhtm/ATH/realtime/api"
	"github.com/dhzjfhtm/ATH/record"
	"github.com/dhzjfhtm/ATH/strategy"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	logger := record.NewLogger()
	binanceFuture := api.NewBinanceFuture(logger)
	//
	config.SetFutureConfig(binanceFuture, logger)

	startFlag := false

	for {
		now := time.Now()
		if now.Hour() == 8 && now.Minute() == 55 {

			// position 조회 -> 0보다 크면 전량 매도 (시장가로)

		} else if now.Hour() == 9 && now.Minute() == 0 {
			// 잔고조회 -> usdt 수량 확인
			// usdtTrade := usdt / len(config.Symbols)
			usdtTrade := 100.

			for i := 0; i < len(config.Symbols); i++ {
				symbol := config.Symbols[i]
				// set leverage and isolated
				binanceFuture.SetLeverage(symbol, 1)
				binanceFuture.SetMarginType(symbol, futures.MarginTypeIsolated)

				// preprocessing
				// 이 부분 함수화
				kline, _ := binanceFuture.GetBinanceFutureKlines(symbol, "d", 2)
				high, _ := strconv.ParseFloat(kline[1].High, 64)
				low, _ := strconv.ParseFloat(kline[1].Low, 64)
				TR := high - low
				open, _ := strconv.ParseFloat(kline[0].Open, 64)
				targetPrice := open + TR*0.5
				// 리턴값이 targetPrice

				// 스트링 변환 -> NewBinanceFutureOrder 함수 내부로 quantity, price parameter는 int로
				quantity := fmt.Sprintf("%f", usdtTrade/targetPrice)
				price := fmt.Sprintf("%f", targetPrice)
				// stop limit order
				binanceFuture.NewBinanceFutureOrder(symbol, "BUY", "STOPLIMIT", quantity, price)
			}

		}
		if startFlag {
			strategy.Larry(binanceFuture)
			time.Sleep(1 * time.Second)
		}

	}

}
