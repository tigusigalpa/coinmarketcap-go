package api

import "github.com/tigusigalpa/coinmarketcap-go/http"

// GlobalMetrics provides global cryptocurrency market data methods.
type GlobalMetrics struct {
	httpClient *http.Client
}

// NewGlobalMetrics creates a global metrics API service.
func NewGlobalMetrics(httpClient *http.Client) *GlobalMetrics {
	return &GlobalMetrics{
		httpClient: httpClient,
	}
}

// QuotesLatest returns the latest global market metrics.
func (g *GlobalMetrics) QuotesLatest(params map[string]string) (map[string]interface{}, error) {
	return g.httpClient.Get("/v1/global-metrics/quotes/latest", params)
}

// QuotesHistorical returns historical global market metrics.
func (g *GlobalMetrics) QuotesHistorical(params map[string]string) (map[string]interface{}, error) {
	return g.httpClient.Get("/v1/global-metrics/quotes/historical", params)
}

// FearAndGreedLatest returns the latest Crypto Fear and Greed Index.
func (g *GlobalMetrics) FearAndGreedLatest(params map[string]string) (map[string]interface{}, error) {
	return g.httpClient.Get("/v3/fear-and-greed/latest", params)
}

// FearAndGreedHistorical returns historical Crypto Fear and Greed Index data.
func (g *GlobalMetrics) FearAndGreedHistorical(params map[string]string) (map[string]interface{}, error) {
	return g.httpClient.Get("/v3/fear-and-greed/historical", params)
}

// AltcoinSeasonIndexLatest returns the latest Altcoin Season Index.
func (g *GlobalMetrics) AltcoinSeasonIndexLatest(params map[string]string) (map[string]interface{}, error) {
	return g.httpClient.Get("/v1/altcoin-season-index/latest", params)
}

// AltcoinSeasonIndexHistorical returns historical Altcoin Season Index data.
func (g *GlobalMetrics) AltcoinSeasonIndexHistorical(params map[string]string) (map[string]interface{}, error) {
	return g.httpClient.Get("/v1/altcoin-season-index/historical", params)
}
