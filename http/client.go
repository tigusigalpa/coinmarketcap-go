package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	cmcerrors "github.com/tigusigalpa/coinmarketcap-go/errors"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	timeout    time.Duration
}

func NewClient(baseURL, apiKey string, timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: strings.TrimRight(baseURL, "/"),
		apiKey:  apiKey,
		timeout: timeout,
	}
}

func (c *Client) Get(endpoint string, params map[string]string) (map[string]interface{}, error) {
	return c.GetWithContext(context.Background(), endpoint, params)
}

// GetWithContext sends a GET request using ctx.
func (c *Client) GetWithContext(ctx context.Context, endpoint string, params map[string]string) (map[string]interface{}, error) {
	fullURL := c.buildURL(endpoint, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, cmcerrors.NewAPIError(fmt.Sprintf("failed to create request: %v", err), 0, nil, err)
	}

	req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, cmcerrors.NewAPIError(fmt.Sprintf("HTTP request failed: %v", err), 0, nil, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, cmcerrors.NewAPIError(fmt.Sprintf("failed to read response body: %v", err), resp.StatusCode, nil, err)
	}

	if resp.StatusCode >= 400 {
		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, cmcerrors.NewAPIError(http.StatusText(resp.StatusCode), resp.StatusCode, nil, err)
		}
		return nil, c.handleError(resp.StatusCode, data)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, cmcerrors.NewAPIError(fmt.Sprintf("failed to parse JSON response: %v", err), resp.StatusCode, nil, err)
	}

	return data, nil
}

func (c *Client) Post(endpoint string, data map[string]interface{}) (map[string]interface{}, error) {
	return c.PostWithContext(context.Background(), endpoint, data)
}

// PostWithContext sends a POST request using ctx.
func (c *Client) PostWithContext(ctx context.Context, endpoint string, data map[string]interface{}) (map[string]interface{}, error) {
	fullURL := c.buildURL(endpoint, nil)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, cmcerrors.NewAPIError(fmt.Sprintf("failed to marshal request data: %v", err), 0, nil, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, cmcerrors.NewAPIError(fmt.Sprintf("failed to create request: %v", err), 0, nil, err)
	}

	req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, cmcerrors.NewAPIError(fmt.Sprintf("HTTP request failed: %v", err), 0, nil, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, cmcerrors.NewAPIError(fmt.Sprintf("failed to read response body: %v", err), resp.StatusCode, nil, err)
	}

	if resp.StatusCode >= 400 {
		var responseData map[string]interface{}
		if err := json.Unmarshal(body, &responseData); err != nil {
			return nil, cmcerrors.NewAPIError(http.StatusText(resp.StatusCode), resp.StatusCode, nil, err)
		}
		return nil, c.handleError(resp.StatusCode, responseData)
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, cmcerrors.NewAPIError(fmt.Sprintf("failed to parse JSON response: %v", err), resp.StatusCode, nil, err)
	}

	return responseData, nil
}

func (c *Client) buildURL(endpoint string, params map[string]string) string {
	fullURL := c.baseURL + "/" + strings.TrimLeft(endpoint, "/")

	if len(params) > 0 {
		values := url.Values{}
		for key, value := range params {
			values.Add(key, value)
		}
		fullURL += "?" + values.Encode()
	}

	return fullURL
}

func (c *Client) handleError(statusCode int, response map[string]interface{}) error {
	message := "Unknown error"
	if status, ok := response["status"].(map[string]interface{}); ok {
		if errorMsg, ok := status["error_message"].(string); ok {
			message = errorMsg
		}
	}

	switch statusCode {
	case 400:
		return cmcerrors.NewInvalidRequestError(message, response)
	case 401:
		return cmcerrors.NewAuthenticationError(message, response)
	case 404:
		return cmcerrors.NewNotFoundError(message, response)
	case 429:
		var retryAfter *int
		if status, ok := response["status"].(map[string]interface{}); ok {
			if ra, ok := status["retry_after"].(float64); ok {
				raInt := int(ra)
				retryAfter = &raInt
			}
		}
		return cmcerrors.NewRateLimitError(message, retryAfter, response)
	default:
		return cmcerrors.NewAPIError(message, statusCode, response, nil)
	}
}
