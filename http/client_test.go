package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	cmcerrors "github.com/tigusigalpa/coinmarketcap-go/errors"
)

func TestGetSendsRequestAndDecodesResponse(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/cryptocurrency/listings/latest" {
			t.Errorf("path = %s", r.URL.Path)
		}
		if r.URL.Query().Get("convert") != "USD,EUR" {
			t.Errorf("convert = %q", r.URL.Query().Get("convert"))
		}
		if r.Header.Get("X-CMC_PRO_API_KEY") != "test-key" {
			t.Errorf("API key header is missing")
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Accept = %q", r.Header.Get("Accept"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"id":1}}`))
	}))
	defer server.Close()

	client := NewClient(server.URL+"/", "test-key", 0)
	data, err := client.Get("v1/cryptocurrency/listings/latest", map[string]string{"convert": "USD,EUR"})
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if data["data"].(map[string]interface{})["id"] != float64(1) {
		t.Fatalf("data = %#v", data)
	}
}

func TestPostWithContextSendsJSON(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/test" {
			t.Errorf("path = %s", r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type = %q", r.Header.Get("Content-Type"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":"ok"}`))
	}))
	defer server.Close()

	data, err := NewClient(server.URL+"/", "test-key", 0).PostWithContext(context.Background(), "v1/test", map[string]interface{}{"id": 1})
	if err != nil {
		t.Fatalf("PostWithContext() error = %v", err)
	}
	if data["data"] != "ok" {
		t.Fatalf("data = %#v", data)
	}
}

func TestGetMapsAPIAndNonJSONErrors(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/rate-limit":
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"status":{"error_message":"Too many requests","retry_after":15}}`))
		default:
			w.WriteHeader(http.StatusBadGateway)
			_, _ = w.Write([]byte("upstream unavailable"))
		}
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", 0)
	_, err := client.Get("/rate-limit", nil)
	var rateLimitErr *cmcerrors.RateLimitError
	if !errors.As(err, &rateLimitErr) || rateLimitErr.RetryAfter == nil || *rateLimitErr.RetryAfter != 15 {
		t.Fatalf("error = %#v, want RateLimitError with retry_after", err)
	}

	_, err = client.Get("/gateway", nil)
	var apiErr *cmcerrors.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error = %v, want APIError", err)
	}
	if apiErr.StatusCode != http.StatusBadGateway || apiErr.Message != http.StatusText(http.StatusBadGateway) {
		t.Fatalf("APIError = %#v", apiErr)
	}
}

func TestBuildURLEncodesParams(t *testing.T) {
	t.Parallel()

	client := NewClient("https://example.test/api/", "test-key", 0)
	parsed, err := url.Parse(client.buildURL("/v1/test", map[string]string{"value": "a value&more"}))
	if err != nil {
		t.Fatalf("url.Parse() error = %v", err)
	}
	if parsed.Path != "/api/v1/test" || parsed.Query().Get("value") != "a value&more" {
		t.Fatalf("URL = %q", parsed.String())
	}
}
