package api

import "github.com/tigusigalpa/coinmarketcap-go/http"

// Exchange provides exchange-related API methods.
type Exchange struct {
	httpClient *http.Client
}

// NewExchange creates an exchange API service.
func NewExchange(httpClient *http.Client) *Exchange {
	return &Exchange{
		httpClient: httpClient,
	}
}

// ListingsLatest returns the latest exchange listings.
func (e *Exchange) ListingsLatest(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/listings/latest", params)
}

// QuotesLatest returns the latest exchange quotes.
func (e *Exchange) QuotesLatest(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/quotes/latest", params)
}

// QuotesHistorical returns historical exchange quotes.
func (e *Exchange) QuotesHistorical(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/quotes/historical", params)
}

// Info returns exchange metadata.
func (e *Exchange) Info(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/info", params)
}

// Map returns CoinMarketCap exchange ID mappings.
func (e *Exchange) Map(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/map", params)
}

// MarketPairsLatest returns the latest exchange market pairs.
func (e *Exchange) MarketPairsLatest(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/market-pairs/latest", params)
}

// Assets returns the reported assets held by an exchange.
func (e *Exchange) Assets(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/assets", params)
}
