package coinmarketcap

import (
	"sync"
	"time"

	"github.com/tigusigalpa/coinmarketcap-go/api"
	"github.com/tigusigalpa/coinmarketcap-go/http"
)

// Client provides access to CoinMarketCap API services.
type Client struct {
	httpClient         *http.Client
	cryptocurrency     *api.Cryptocurrency
	cryptocurrencyOnce sync.Once
	exchange           *api.Exchange
	exchangeOnce       sync.Once
	globalMetrics      *api.GlobalMetrics
	globalMetricsOnce  sync.Once
	tools              *api.Tools
	toolsOnce          sync.Once
}

// NewClient creates a client backed by httpClient.
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Cryptocurrency returns the cryptocurrency API service.
func (c *Client) Cryptocurrency() *api.Cryptocurrency {
	c.cryptocurrencyOnce.Do(func() {
		c.cryptocurrency = api.NewCryptocurrency(c.httpClient)
	})
	return c.cryptocurrency
}

// Exchange returns the exchange API service.
func (c *Client) Exchange() *api.Exchange {
	c.exchangeOnce.Do(func() {
		c.exchange = api.NewExchange(c.httpClient)
	})
	return c.exchange
}

// GlobalMetrics returns the global market metrics API service.
func (c *Client) GlobalMetrics() *api.GlobalMetrics {
	c.globalMetricsOnce.Do(func() {
		c.globalMetrics = api.NewGlobalMetrics(c.httpClient)
	})
	return c.globalMetrics
}

// Tools returns the utility API service.
func (c *Client) Tools() *api.Tools {
	c.toolsOnce.Do(func() {
		c.tools = api.NewTools(c.httpClient)
	})
	return c.tools
}

// ClientBuilder configures and creates a Client.
type ClientBuilder struct {
	apiKey  string
	baseURL string
	timeout time.Duration
}

// NewClientBuilder creates a builder with production defaults.
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		baseURL: "https://pro-api.coinmarketcap.com",
		timeout: 30 * time.Second,
	}
}

// SetAPIKey sets the CoinMarketCap API key.
func (b *ClientBuilder) SetAPIKey(apiKey string) *ClientBuilder {
	b.apiKey = apiKey
	return b
}

// SetBaseURL sets the API base URL.
func (b *ClientBuilder) SetBaseURL(baseURL string) *ClientBuilder {
	b.baseURL = baseURL
	return b
}

// SetTimeout sets the HTTP request timeout.
func (b *ClientBuilder) SetTimeout(timeout time.Duration) *ClientBuilder {
	b.timeout = timeout
	return b
}

// UseSandbox configures the client for the CoinMarketCap sandbox API.
func (b *ClientBuilder) UseSandbox() *ClientBuilder {
	b.baseURL = "https://sandbox-api.coinmarketcap.com"
	return b
}

// Build validates the configuration and creates a Client.
func (b *ClientBuilder) Build() (*Client, error) {
	if b.apiKey == "" {
		return nil, &ClientBuilderError{Message: "API key is required"}
	}

	httpClient := http.NewClient(b.baseURL, b.apiKey, b.timeout)
	return NewClient(httpClient), nil
}

// ClientBuilderError describes an invalid client builder configuration.
type ClientBuilderError struct {
	Message string
}

func (e *ClientBuilderError) Error() string {
	return e.Message
}
