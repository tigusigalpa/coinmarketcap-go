package coinmarketcap

import (
	"sync"
	"time"

	"github.com/tigusigalpa/coinmarketcap-go/api"
	"github.com/tigusigalpa/coinmarketcap-go/http"
)

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

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) Cryptocurrency() *api.Cryptocurrency {
	c.cryptocurrencyOnce.Do(func() {
		c.cryptocurrency = api.NewCryptocurrency(c.httpClient)
	})
	return c.cryptocurrency
}

func (c *Client) Exchange() *api.Exchange {
	c.exchangeOnce.Do(func() {
		c.exchange = api.NewExchange(c.httpClient)
	})
	return c.exchange
}

func (c *Client) GlobalMetrics() *api.GlobalMetrics {
	c.globalMetricsOnce.Do(func() {
		c.globalMetrics = api.NewGlobalMetrics(c.httpClient)
	})
	return c.globalMetrics
}

func (c *Client) Tools() *api.Tools {
	c.toolsOnce.Do(func() {
		c.tools = api.NewTools(c.httpClient)
	})
	return c.tools
}

type ClientBuilder struct {
	apiKey  string
	baseURL string
	timeout time.Duration
}

func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		baseURL: "https://pro-api.coinmarketcap.com",
		timeout: 30 * time.Second,
	}
}

func (b *ClientBuilder) SetAPIKey(apiKey string) *ClientBuilder {
	b.apiKey = apiKey
	return b
}

func (b *ClientBuilder) SetBaseURL(baseURL string) *ClientBuilder {
	b.baseURL = baseURL
	return b
}

func (b *ClientBuilder) SetTimeout(timeout time.Duration) *ClientBuilder {
	b.timeout = timeout
	return b
}

func (b *ClientBuilder) UseSandbox() *ClientBuilder {
	b.baseURL = "https://sandbox-api.coinmarketcap.com"
	return b
}

func (b *ClientBuilder) Build() (*Client, error) {
	if b.apiKey == "" {
		return nil, &ClientBuilderError{Message: "API key is required"}
	}

	httpClient := http.NewClient(b.baseURL, b.apiKey, b.timeout)
	return NewClient(httpClient), nil
}

type ClientBuilderError struct {
	Message string
}

func (e *ClientBuilderError) Error() string {
	return e.Message
}
