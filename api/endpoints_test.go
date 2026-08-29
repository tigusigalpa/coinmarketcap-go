package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	cmchttp "github.com/tigusigalpa/coinmarketcap-go/http"
)

func TestNewEndpointsUseDocumentedPaths(t *testing.T) {
	t.Parallel()

	paths := make(chan string, 20)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths <- r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{}}`))
	}))
	defer server.Close()

	client := cmchttp.NewClient(server.URL, "test-key", time.Second)
	crypto := NewCryptocurrency(client)
	exchange := NewExchange(client)

	calls := []func(map[string]string) (map[string]interface{}, error){
		crypto.ListingsLatest,
		crypto.QuotesLatest,
		crypto.ListingsNew,
		crypto.Category,
		crypto.MarketPairsLatest,
		crypto.PricePerformanceStatsLatest,
		crypto.TrendingMostVisited,
		crypto.QuotesHistorical,
		crypto.SimplePrice,
		crypto.Airdrops,
		crypto.Airdrop,
		exchange.Assets,
		NewGlobalMetrics(client).FearAndGreedLatest,
		NewGlobalMetrics(client).FearAndGreedHistorical,
		NewGlobalMetrics(client).AltcoinSeasonIndexLatest,
		NewGlobalMetrics(client).AltcoinSeasonIndexHistorical,
	}
	want := []string{
		"/v3/cryptocurrency/listings/latest",
		"/v3/cryptocurrency/quotes/latest",
		"/v1/cryptocurrency/listings/new",
		"/v1/cryptocurrency/category",
		"/v2/cryptocurrency/market-pairs/latest",
		"/v2/cryptocurrency/price-performance-stats/latest",
		"/v1/cryptocurrency/trending/most-visited",
		"/v3/cryptocurrency/quotes/historical",
		"/v1/simple/price",
		"/v1/cryptocurrency/airdrops",
		"/v1/cryptocurrency/airdrop",
		"/v1/exchange/assets",
		"/v3/fear-and-greed/latest",
		"/v3/fear-and-greed/historical",
		"/v1/altcoin-season-index/latest",
		"/v1/altcoin-season-index/historical",
	}

	for i, call := range calls {
		if _, err := call(map[string]string{"id": "1"}); err != nil {
			t.Fatalf("call %d returned error: %v", i, err)
		}
		if got := <-paths; got != want[i] {
			t.Fatalf("call %d path = %q, want %q", i, got, want[i])
		}
	}
}
