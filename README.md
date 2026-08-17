# CoinMarketCap API Client for Go

![CoinMarketCap Go SDK](https://i.postimg.cc/tTfkMhvb/coinmarketcap-golang-api.jpg)

**🌐 Language:** English | [Русский](README-ru.md)

**📦 Package:** [coinmarketcap package](https://pkg.go.dev/github.com/tigusigalpa/coinmarketcap-go)

**Professional CoinMarketCap API v1 Integration for Golang Applications**

A modern, production-ready Go package that provides seamless integration with the CoinMarketCap cryptocurrency API.
Built for Go developers who need reliable access to real-time cryptocurrency prices, market data, exchange information,
and blockchain analytics.

## 🌟 Why Choose This CoinMarketCap Go Client?

Whether you're building a cryptocurrency portfolio tracker, a crypto trading bot, a blockchain analytics dashboard, or
integrating real-time Bitcoin and Ethereum prices into your Go application, this package provides everything you need:

- **Production-Ready**: Battle-tested code with comprehensive error handling for cryptocurrency API integration
- **Developer-Friendly**: Intuitive fluent interface designed specifically for Golang developers
- **Type-Safe**: Full Go type safety with structured error handling
- **Well-Documented**: Extensive documentation with real-world cryptocurrency integration examples
- **Performance Optimized**: Efficient HTTP client implementation for fast cryptocurrency data retrieval
- **Flexible Architecture**: Easy to integrate into any Go application

## 📊 What Can You Build?

This CoinMarketCap API wrapper enables you to create:

- **Cryptocurrency Portfolio Trackers** - Track Bitcoin, Ethereum, and 10,000+ cryptocurrencies in real-time
- **Crypto Price Alert Systems** - Monitor cryptocurrency prices and send notifications
- **Trading Dashboards** - Display live crypto market data, charts, and trading volumes
- **Blockchain Analytics Platforms** - Analyze cryptocurrency market trends and dominance
- **DeFi Applications** - Integrate cryptocurrency price feeds for DeFi protocols
- **Crypto News Portals** - Show real-time cryptocurrency prices alongside news
- **Exchange Comparison Tools** - Compare cryptocurrency prices across different exchanges
- **Market Research Tools** - Access historical cryptocurrency data for analysis
- **Crypto Converter Apps** - Convert between different cryptocurrencies and fiat currencies
- **Investment Analysis Tools** - Track cryptocurrency market capitalization and performance

## ✨ Key Features

### Core Functionality

- 🚀 **Modern Go 1.21+** - Leverages latest Go features
- 🎯 **Standard HTTP Client** - Built on Go's standard `net/http` package
- 🏗️ **Fluent ClientBuilder** - Elegant, chainable API for easy client configuration
- 🔒 **Type-safe Error Handling** - Strongly-typed error types for all API errors
- ⚡ **Smart Exception Handling** - Custom errors for rate limits, authentication, and API errors
- 🧪 **Easy to Test** - Clean architecture for unit testing
- 📊 **Comprehensive Coverage** - All CoinMarketCap API v1 endpoints supported

### Cryptocurrency API Coverage

- 📈 **Real-time Cryptocurrency Prices** - Get live Bitcoin, Ethereum, and altcoin prices
- 💰 **Market Capitalization Data** - Access market cap rankings for 10,000+ cryptocurrencies
- 📊 **Historical Price Data** - Retrieve historical cryptocurrency quotes and OHLCV data
- 🔄 **Price Conversion Tools** - Convert between cryptocurrencies and fiat currencies
- 🏆 **Trending Cryptocurrencies** - Track trending coins, gainers, and losers
- 🏷️ **Cryptocurrency Metadata** - Get logos, descriptions, websites, and social links
- 📉 **OHLCV Data** - Access Open, High, Low, Close, Volume data for technical analysis

### Exchange & Market Data

- 🏦 **Exchange Information** - Data for 300+ cryptocurrency exchanges
- 💱 **Trading Pairs** - Access market pairs and trading volume data
- 🌍 **Global Market Metrics** - Total market cap, Bitcoin dominance, and market statistics
- 📊 **Exchange Rankings** - Sort exchanges by volume, liquidity, and other metrics

## Requirements

- Golang 1.21 or higher
- CoinMarketCap API key ([Get one here](https://coinmarketcap.com/api/))

## 📦 Installation Guide

### Step 1: Install via Go Get

Install the CoinMarketCap API client package using Go modules:

```bash
go get github.com/tigusigalpa/coinmarketcap-go
```

### Step 2: Get Your CoinMarketCap API Key

Before you can use this package, you need a CoinMarketCap API key:

1. Visit [CoinMarketCap API Portal](https://coinmarketcap.com/api/)
2. Sign up for a free account (Basic plan includes 10,000 API calls/month)
3. Navigate to your API dashboard
4. Copy your API key

**API Plans Available:**

- **Basic (Free)** - 10,000 calls/month, perfect for development and small projects
- **Hobbyist** - 30,000 calls/month for personal crypto projects
- **Startup** - 100,000 calls/month for growing applications
- **Standard** - 300,000 calls/month for production applications
- **Professional** - 1,000,000+ calls/month for enterprise solutions

### Step 3: Configure Environment Variables (Optional)

Create a `.env` file in your project root:

```env
COINMARKETCAP_API_KEY=your-api-key-here
COINMARKETCAP_BASE_URL=https://pro-api.coinmarketcap.com
COINMARKETCAP_TIMEOUT=30
COINMARKETCAP_USE_SANDBOX=false
```

### Step 4: Import and Use

```go
package main

import (
    "fmt"
    "log"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("your-api-key-here").
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    // Get Bitcoin price
    result, err := client.Cryptocurrency().QuotesLatest(map[string]string{
        "symbol":  "BTC",
        "convert": "USD",
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Bitcoin Price: %v\n", result)
}
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
    // Create client with builder pattern
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("your-api-key").
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    // Get top 10 cryptocurrencies
    listings, err := client.Cryptocurrency().ListingsLatest(map[string]string{
        "start":   "1",
        "limit":   "10",
        "convert": "USD",
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Top 10 Cryptocurrencies: %v\n", listings)
}
```

### Using Custom Configuration

```go
client, err := coinmarketcap.NewClientBuilder().
    SetAPIKey("your-api-key").
    SetBaseURL("https://pro-api.coinmarketcap.com").
    SetTimeout(60 * time.Second).
    Build()
```

### Using Sandbox Mode

```go
client, err := coinmarketcap.NewClientBuilder().
    SetAPIKey("your-sandbox-api-key").
    UseSandbox().
    Build()
```

## Available API Methods

### Cryptocurrency API

| Method                    | Endpoint                                     | Description                                  |
|---------------------------|----------------------------------------------|----------------------------------------------|
| `ListingsLatest()`        | `/v1/cryptocurrency/listings/latest`         | Get list of cryptocurrencies by market cap   |
| `ListingsHistorical()`    | `/v1/cryptocurrency/listings/historical`     | Historical listings                          |
| `QuotesLatest()`          | `/v2/cryptocurrency/quotes/latest`           | Current quotes for specific cryptocurrencies |
| `QuotesHistorical()`      | `/v2/cryptocurrency/quotes/historical`       | Historical quotes                            |
| `Info()`                  | `/v2/cryptocurrency/info`                    | Metadata (logos, links, description)         |
| `Map()`                   | `/v1/cryptocurrency/map`                     | CoinMarketCap ID mapping                     |
| `OHLCVLatest()`           | `/v2/cryptocurrency/ohlcv/latest`            | Latest OHLCV data                            |
| `OHLCVHistorical()`       | `/v2/cryptocurrency/ohlcv/historical`        | Historical OHLCV data                        |
| `Categories()`            | `/v1/cryptocurrency/categories`              | Cryptocurrency categories                    |
| `TrendingLatest()`        | `/v1/cryptocurrency/trending/latest`         | Trending cryptocurrencies                    |
| `TrendingGainersLosers()` | `/v1/cryptocurrency/trending/gainers-losers` | Top gainers and losers                       |

### Exchange API

| Method                | Endpoint                           | Description                         |
|-----------------------|------------------------------------|-------------------------------------|
| `ListingsLatest()`    | `/v1/exchange/listings/latest`     | List of exchanges by trading volume |
| `QuotesLatest()`      | `/v1/exchange/quotes/latest`       | Current exchange data               |
| `QuotesHistorical()`  | `/v1/exchange/quotes/historical`   | Historical data                     |
| `Info()`              | `/v1/exchange/info`                | Exchange metadata                   |
| `Map()`               | `/v1/exchange/map`                 | Exchange ID mapping                 |
| `MarketPairsLatest()` | `/v1/exchange/market-pairs/latest` | Exchange trading pairs              |

### Global Metrics API

| Method               | Endpoint                               | Description               |
|----------------------|----------------------------------------|---------------------------|
| `QuotesLatest()`     | `/v1/global-metrics/quotes/latest`     | Current global metrics    |
| `QuotesHistorical()` | `/v1/global-metrics/quotes/historical` | Historical global metrics |

### Tools API

| Method              | Endpoint                     | Description                       |
|---------------------|------------------------------|-----------------------------------|
| `PriceConversion()` | `/v2/tools/price-conversion` | Convert prices between currencies |

## 💡 Real-World Usage Examples

### Example 1: Get Bitcoin and Ethereum Prices

```go
package main

import (
    "fmt"
    "log"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("your-api-key").
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    quotes, err := client.Cryptocurrency().QuotesLatest(map[string]string{
        "symbol":  "BTC,ETH",
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
    
    eth := data["ETH"].(map[string]interface{})
    ethQuote := eth["quote"].(map[string]interface{})
    ethUSD := ethQuote["USD"].(map[string]interface{})
    
    fmt.Printf("Ethereum Price: $%.2f\n", ethUSD["price"].(float64))
}
```

### Example 2: Track Top 100 Cryptocurrencies

```go
package main

import (
    "fmt"
    "log"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("your-api-key").
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    listings, err := client.Cryptocurrency().ListingsLatest(map[string]string{
        "start":    "1",
        "limit":    "100",
        "convert":  "USD",
        "sort":     "market_cap",
        "sort_dir": "desc",
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    data := listings["data"].([]interface{})
    
    for _, item := range data {
        crypto := item.(map[string]interface{})
        quote := crypto["quote"].(map[string]interface{})
        usd := quote["USD"].(map[string]interface{})
        
        fmt.Printf("%d. %s (%s) - $%.2f\n",
            int(crypto["cmc_rank"].(float64)),
            crypto["name"].(string),
            crypto["symbol"].(string),
            usd["price"].(float64),
        )
    }
}
```

### Example 3: Price Conversion

```go
package main

import (
    "fmt"
    "log"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("your-api-key").
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    conversion, err := client.Tools().PriceConversion(map[string]string{
        "amount":  "1",
        "symbol":  "BTC",
        "convert": "ETH",
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Conversion Result: %v\n", conversion)
}
```

### Example 4: Get Global Market Metrics

```go
package main

import (
    "fmt"
    "log"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("your-api-key").
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    metrics, err := client.GlobalMetrics().QuotesLatest(map[string]string{
        "convert": "USD",
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    data := metrics["data"].(map[string]interface{})
    quote := data["quote"].(map[string]interface{})
    usd := quote["USD"].(map[string]interface{})
    
    fmt.Printf("Total Market Cap: $%.2f\n", usd["total_market_cap"].(float64))
    fmt.Printf("Total Volume 24h: $%.2f\n", usd["total_volume_24h"].(float64))
    fmt.Printf("BTC Dominance: %.2f%%\n", data["btc_dominance"].(float64))
}
```

## 🛡️ Error Handling

The package provides comprehensive error handling for all CoinMarketCap API errors:

### Error Types

| Error Type            | HTTP Code | Description                |
|-----------------------|-----------|----------------------------|
| `AuthenticationError` | 401       | Invalid or missing API key |
| `RateLimitError`      | 429       | API rate limit exceeded    |
| `InvalidRequestError` | 400       | Invalid request parameters |
| `NotFoundError`       | 404       | Resource not found         |
| `APIError`            | Various   | General API error          |

### Basic Error Handling

```go
package main

import (
    "errors"
    "fmt"
    "log"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
    cmcerrors "github.com/tigusigalpa/coinmarketcap-go/errors"
)

func main() {
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("your-api-key").
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    listings, err := client.Cryptocurrency().ListingsLatest(map[string]string{
        "limit":   "100",
        "convert": "USD",
    })
    
    if err != nil {
        var authErr *cmcerrors.AuthenticationError
        var rateLimitErr *cmcerrors.RateLimitError
        var invalidReqErr *cmcerrors.InvalidRequestError
        var notFoundErr *cmcerrors.NotFoundError
        var apiErr *cmcerrors.APIError
        
        switch {
        case errors.As(err, &authErr):
            fmt.Printf("Authentication failed: %v\n", authErr)
        case errors.As(err, &rateLimitErr):
            fmt.Printf("Rate limit exceeded. Retry after: %v\n", rateLimitErr.GetRetryAfter())
        case errors.As(err, &invalidReqErr):
            fmt.Printf("Invalid request: %v\n", invalidReqErr)
        case errors.As(err, &notFoundErr):
            fmt.Printf("Resource not found: %v\n", notFoundErr)
        case errors.As(err, &apiErr):
            fmt.Printf("API error [%d]: %v\n", apiErr.StatusCode, apiErr)
        default:
            fmt.Printf("Unknown error: %v\n", err)
        }
        return
    }
    
    fmt.Printf("Listings: %v\n", listings)
}
```

### Retry Logic for Rate Limits

```go
package main

import (
    "errors"
    "fmt"
    "log"
    "time"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
    cmcerrors "github.com/tigusigalpa/coinmarketcap-go/errors"
)

func getCryptoPricesWithRetry(client *coinmarketcap.Client, symbols string, maxRetries int) (map[string]interface{}, error) {
    for attempt := 0; attempt < maxRetries; attempt++ {
        result, err := client.Cryptocurrency().QuotesLatest(map[string]string{
            "symbol":  symbols,
            "convert": "USD",
        })
        
        if err == nil {
            return result, nil
        }
        
        var rateLimitErr *cmcerrors.RateLimitError
        if errors.As(err, &rateLimitErr) {
            if attempt >= maxRetries-1 {
                return nil, err
            }
            
            retryAfter := 60
            if rateLimitErr.GetRetryAfter() != nil {
                retryAfter = *rateLimitErr.GetRetryAfter()
            }
            
            fmt.Printf("Rate limit hit. Retrying after %d seconds...\n", retryAfter)
            time.Sleep(time.Duration(retryAfter) * time.Second)
            continue
        }
        
        return nil, err
    }
    
    return nil, fmt.Errorf("max retries exceeded")
}

func main() {
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("your-api-key").
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    result, err := getCryptoPricesWithRetry(client, "BTC,ETH", 3)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Result: %v\n", result)
}
```

## 🧪 Testing

### Running Tests

```bash
go test ./...
```

### Running Tests with Coverage

```bash
go test -cover ./...
```

## 📝 License

MIT License - see LICENSE file for details

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📧 Support

For issues, questions, or contributions, please visit
the [GitHub repository](https://github.com/tigusigalpa/coinmarketcap-go).

## 🔗 Links

- [CoinMarketCap API Documentation](https://coinmarketcap.com/api/documentation/v1/)
- [Get API Key](https://coinmarketcap.com/api/)
- [GitHub Repository](https://github.com/tigusigalpa/coinmarketcap-go)
