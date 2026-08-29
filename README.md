# CoinMarketCap client for Go

[![CI](https://github.com/tigusigalpa/coinmarketcap-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/coinmarketcap-go/actions/workflows/ci.yml)
[![Tests](https://github.com/tigusigalpa/coinmarketcap-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/coinmarketcap-go/actions/workflows/test.yml)
[![CodeQL](https://github.com/tigusigalpa/coinmarketcap-go/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/coinmarketcap-go/actions/workflows/codeql.yml)
[![Codecov](https://codecov.io/gh/tigusigalpa/coinmarketcap-go/graph/badge.svg)](https://codecov.io/gh/tigusigalpa/coinmarketcap-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/tigusigalpa/coinmarketcap-go.svg)](https://pkg.go.dev/github.com/tigusigalpa/coinmarketcap-go)
[![Go version](https://img.shields.io/github/go-mod/go-version/tigusigalpa/coinmarketcap-go)](go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigusigalpa/coinmarketcap-go)](https://goreportcard.com/report/github.com/tigusigalpa/coinmarketcap-go)
[![License](https://img.shields.io/github/license/tigusigalpa/coinmarketcap-go)](LICENSE)

![CoinMarketCap Go SDK](https://i.postimg.cc/tTfkMhvb/coinmarketcap-golang-api.jpg)

**Language:** English · [Русский](README-ru.md)
**Package:** [pkg.go.dev/github.com/tigusigalpa/coinmarketcap-go](https://pkg.go.dev/github.com/tigusigalpa/coinmarketcap-go)

A small Go client for the [CoinMarketCap API](https://coinmarketcap.com/api/documentation/v1/). Use one configured client to retrieve cryptocurrency quotes and listings, exchange data, market-wide metrics, and price conversions.

Responses are returned as `map[string]interface{}` to stay close to CoinMarketCap's JSON. This means new query parameters work without waiting for model updates. If your application uses structs, convert the returned `data` to the shape you need.

## Requirements

- Go 1.21 or newer
- A CoinMarketCap API key — [create one here](https://coinmarketcap.com/api/)

## Install

```bash
go get github.com/tigusigalpa/coinmarketcap-go
```

## Quick start

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

	fmt.Printf("%#v\n", quotes["data"])
}
```

Every method accepts a `map[string]string` of query parameters. Refer to the [official API documentation](https://coinmarketcap.com/api/documentation/v1/) for endpoint-specific parameters, response shapes, and plan availability.

## Configuration

The default target is the production API, with a 30-second timeout.

```go
client, err := coinmarketcap.NewClientBuilder().
	SetAPIKey("your-api-key").
	SetTimeout(60 * time.Second).
	Build()
```

For the sandbox, use a sandbox API key and call `UseSandbox()`:

```go
client, err := coinmarketcap.NewClientBuilder().
	SetAPIKey("your-sandbox-api-key").
	UseSandbox().
	Build()
```

`SetBaseURL()` is useful for another compatible endpoint, such as a local test server.

## Supported endpoints

| Service | Methods |
| --- | --- |
| `Cryptocurrency()` | `ListingsLatest`, `ListingsHistorical`, `ListingsNew`, `QuotesLatest`, `QuotesHistorical`, `Info`, `Map`, `OHLCVLatest`, `OHLCVHistorical`, `Categories`, `Category`, `MarketPairsLatest`, `PricePerformanceStatsLatest`, `SimplePrice`, `Airdrops`, `Airdrop`, `TrendingLatest`, `TrendingGainersLosers`, `TrendingMostVisited` |
| `Exchange()` | `ListingsLatest`, `QuotesLatest`, `QuotesHistorical`, `Info`, `Map`, `MarketPairsLatest`, `Assets` |
| `GlobalMetrics()` | `QuotesLatest`, `QuotesHistorical`, `FearAndGreedLatest`, `FearAndGreedHistorical`, `AltcoinSeasonIndexLatest`, `AltcoinSeasonIndexHistorical` |
| `Tools()` | `PriceConversion` |

For example, fetch the first ten listings:

```go
listings, err := client.Cryptocurrency().ListingsLatest(map[string]string{
	"start": "1", "limit": "10", "convert": "USD",
})
```

## Errors

Use `errors.As` to handle errors you care about:

| Type | HTTP status |
| --- | --- |
| `*errors.AuthenticationError` | 401 |
| `*errors.InvalidRequestError` | 400 |
| `*errors.NotFoundError` | 404 |
| `*errors.RateLimitError` | 429 |
| `*errors.APIError` | Other request or API errors |

```go
var rateLimitErr *cmcerrors.RateLimitError
if errors.As(err, &rateLimitErr) {
	if retryAfter := rateLimitErr.GetRetryAfter(); retryAfter != nil {
		fmt.Printf("Try again in %d seconds\n", *retryAfter)
	}
}
```

`APIError` includes the HTTP status, parsed response when available, and the underlying transport or decoding error.

## Testing

```bash
go test ./...
go test -race ./...
```

Issues and pull requests are welcome. The [official CoinMarketCap documentation](https://coinmarketcap.com/api/documentation/v1/) is the source of truth for API parameters, quotas, and data availability.

## License

[MIT](LICENSE)
