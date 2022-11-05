package main

import (
	"fmt"

	"github.com/dhzjfhtm/ATH/realtime/api"
	"github.com/dhzjfhtm/ATH/record"
	"github.com/dhzjfhtm/ATH/trade"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	logger := record.NewLogger()
	binanceFuture := api.NewBinanceFuture(logger)
	trade.FutureTrade(binanceFuture)
}
