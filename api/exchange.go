package api

import "github.com/tigusigalpa/coinmarketcap-go/http"

type Exchange struct {
	httpClient *http.Client
}

func NewExchange(httpClient *http.Client) *Exchange {
	return &Exchange{
		httpClient: httpClient,
	}
}

func (e *Exchange) ListingsLatest(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/listings/latest", params)
}

func (e *Exchange) QuotesLatest(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/quotes/latest", params)
}

func (e *Exchange) QuotesHistorical(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/quotes/historical", params)
}

func (e *Exchange) Info(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/info", params)
}

func (e *Exchange) Map(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/map", params)
}

func (e *Exchange) MarketPairsLatest(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/market-pairs/latest", params)
}

func (e *Exchange) Assets(params map[string]string) (map[string]interface{}, error) {
	return e.httpClient.Get("/v1/exchange/assets", params)
}
