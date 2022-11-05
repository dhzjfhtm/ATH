package main

import (
	"fmt"

	"github.com/dhzjfhtm/ATH/record"
	"github.com/dhzjfhtm/ATH/trade"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	//trade.SpotTrade()
	trade.FutureTrade()
	// Use Test() instead of Do() for testing.

	logger := record.NewLogger()

	logger.Info("Hello World")
}
