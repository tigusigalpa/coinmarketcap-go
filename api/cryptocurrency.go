package api

import "github.com/tigusigalpa/coinmarketcap-go/http"

// Cryptocurrency provides cryptocurrency-related API methods.
type Cryptocurrency struct {
	httpClient *http.Client
}

// NewCryptocurrency creates a cryptocurrency API service.
func NewCryptocurrency(httpClient *http.Client) *Cryptocurrency {
	return &Cryptocurrency{
		httpClient: httpClient,
	}
}

// ListingsLatest returns the latest cryptocurrency listings.
func (c *Cryptocurrency) ListingsLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v3/cryptocurrency/listings/latest", params)
}

// ListingsHistorical returns a historical cryptocurrency listings snapshot.
func (c *Cryptocurrency) ListingsHistorical(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/listings/historical", params)
}

// ListingsNew returns recently added cryptocurrencies.
func (c *Cryptocurrency) ListingsNew(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/listings/new", params)
}

// QuotesLatest returns the latest cryptocurrency quotes.
func (c *Cryptocurrency) QuotesLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v3/cryptocurrency/quotes/latest", params)
}

// QuotesHistorical returns historical cryptocurrency quotes.
func (c *Cryptocurrency) QuotesHistorical(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v3/cryptocurrency/quotes/historical", params)
}

// Info returns cryptocurrency metadata.
func (c *Cryptocurrency) Info(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/info", params)
}

// Map returns CoinMarketCap cryptocurrency ID mappings.
func (c *Cryptocurrency) Map(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/map", params)
}

// OHLCVLatest returns the latest cryptocurrency OHLCV data.
func (c *Cryptocurrency) OHLCVLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/ohlcv/latest", params)
}

// OHLCVHistorical returns historical cryptocurrency OHLCV data.
func (c *Cryptocurrency) OHLCVHistorical(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/ohlcv/historical", params)
}

// Categories returns cryptocurrency categories.
func (c *Cryptocurrency) Categories(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/categories", params)
}

// Category returns details for one cryptocurrency category.
func (c *Cryptocurrency) Category(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/category", params)
}

// MarketPairsLatest returns the latest cryptocurrency market pairs.
func (c *Cryptocurrency) MarketPairsLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/market-pairs/latest", params)
}

// PricePerformanceStatsLatest returns current price performance statistics.
func (c *Cryptocurrency) PricePerformanceStatsLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/price-performance-stats/latest", params)
}

// SimplePrice returns simplified cryptocurrency prices.
func (c *Cryptocurrency) SimplePrice(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/simple/price", params)
}

// Airdrops returns active and upcoming airdrops.
func (c *Cryptocurrency) Airdrops(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/airdrops", params)
}

// Airdrop returns details for one airdrop.
func (c *Cryptocurrency) Airdrop(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/airdrop", params)
}

// TrendingLatest returns currently trending cryptocurrencies.
func (c *Cryptocurrency) TrendingLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/trending/latest", params)
}

// TrendingGainersLosers returns the top cryptocurrency gainers and losers.
func (c *Cryptocurrency) TrendingGainersLosers(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/trending/gainers-losers", params)
}

// TrendingMostVisited returns the most visited cryptocurrencies on CoinMarketCap.
func (c *Cryptocurrency) TrendingMostVisited(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/trending/most-visited", params)
}
