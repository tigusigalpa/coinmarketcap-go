package api

import "github.com/tigusigalpa/coinmarketcap-go/http"

type Cryptocurrency struct {
	httpClient *http.Client
}

func NewCryptocurrency(httpClient *http.Client) *Cryptocurrency {
	return &Cryptocurrency{
		httpClient: httpClient,
	}
}

func (c *Cryptocurrency) ListingsLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v3/cryptocurrency/listings/latest", params)
}

func (c *Cryptocurrency) ListingsHistorical(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/listings/historical", params)
}

func (c *Cryptocurrency) ListingsNew(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/listings/new", params)
}

func (c *Cryptocurrency) QuotesLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v3/cryptocurrency/quotes/latest", params)
}

func (c *Cryptocurrency) QuotesHistorical(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v3/cryptocurrency/quotes/historical", params)
}

func (c *Cryptocurrency) Info(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/info", params)
}

func (c *Cryptocurrency) Map(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/map", params)
}

func (c *Cryptocurrency) OHLCVLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/ohlcv/latest", params)
}

func (c *Cryptocurrency) OHLCVHistorical(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/ohlcv/historical", params)
}

func (c *Cryptocurrency) Categories(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/categories", params)
}

func (c *Cryptocurrency) Category(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/category", params)
}

func (c *Cryptocurrency) MarketPairsLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/market-pairs/latest", params)
}

func (c *Cryptocurrency) PricePerformanceStatsLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/price-performance-stats/latest", params)
}

func (c *Cryptocurrency) SimplePrice(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/simple/price", params)
}

func (c *Cryptocurrency) Airdrops(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/airdrops", params)
}

func (c *Cryptocurrency) Airdrop(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/airdrop", params)
}

func (c *Cryptocurrency) TrendingLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/trending/latest", params)
}

func (c *Cryptocurrency) TrendingGainersLosers(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/trending/gainers-losers", params)
}

func (c *Cryptocurrency) TrendingMostVisited(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/trending/most-visited", params)
}
