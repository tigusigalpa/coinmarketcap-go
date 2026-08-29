package api

import "github.com/tigusigalpa/coinmarketcap-go/http"

// Tools provides CoinMarketCap utility API methods.
type Tools struct {
	httpClient *http.Client
}

// NewTools creates a utility API service.
func NewTools(httpClient *http.Client) *Tools {
	return &Tools{
		httpClient: httpClient,
	}
}

// PriceConversion converts a cryptocurrency amount to another currency.
func (t *Tools) PriceConversion(params map[string]string) (map[string]interface{}, error) {
	return t.httpClient.Get("/v2/tools/price-conversion", params)
}
