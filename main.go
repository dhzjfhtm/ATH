package main

import (
	"fmt"
	"time"

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
		// 시간 체크 9:00
		now := time.Now()
		if now.Hour() == 8 && now.Minute() == 55 {

			// 청산

		} else if now.Hour() == 9 && now.Minute() == 0 {
			// set leverage and isolated

			// preprocessing
			// range := 전날 high - low
			// target price = 오늘 시가 + k(0.5) * range

			// stop limit order

		}
		if startFlag {
			strategy.Larry(binanceFuture)
			time.Sleep(1 * time.Second)
		}

	}

}
