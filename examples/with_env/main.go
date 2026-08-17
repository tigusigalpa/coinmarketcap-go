package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	apiKey := os.Getenv("COINMARKETCAP_API_KEY")
	if apiKey == "" {
		log.Fatal("COINMARKETCAP_API_KEY environment variable is required")
	}

	builder := coinmarketcap.NewClientBuilder().SetAPIKey(apiKey)

	if baseURL := os.Getenv("COINMARKETCAP_BASE_URL"); baseURL != "" {
		builder = builder.SetBaseURL(baseURL)
	}

	if os.Getenv("COINMARKETCAP_USE_SANDBOX") == "true" {
		builder = builder.UseSandbox()
	}

	client, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	quotes, err := client.Cryptocurrency().QuotesLatest(map[string]string{
		"symbol":  "BTC",
		"convert": "USD",
	})

	if err != nil {
		log.Fatal(err)
	}

	data := quotes["data"].(map[string]interface{})
	btc := data["BTC"].(map[string]interface{})
	btcQuote := btc["quote"].(map[string]interface{})
	btcUSD := btcQuote["USD"].(map[string]interface{})

	fmt.Printf("Bitcoin Price: $%.2f\n", btcUSD["price"].(float64))
	fmt.Printf("24h Change: %.2f%%\n", btcUSD["percent_change_24h"].(float64))
	fmt.Printf("Market Cap: $%.0f\n", btcUSD["market_cap"].(float64))
}
