package api_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
	cmchttp "github.com/tigusigalpa/coinmarketcap-go/http"
)

func TestFearAndGreedLatest(t *testing.T) {
	t.Parallel()

	// The fixture's certificate is trusted by the injected client's transport.
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/fear-and-greed/latest" {
			t.Errorf("request = %s %s, want GET /v3/fear-and-greed/latest", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("X-CMC_PRO_API_KEY"); got != "test-key" {
			t.Errorf("X-CMC_PRO_API_KEY = %q, want test-key", got)
		}
		if got := r.URL.Query().Get("test"); got != "a value&more" {
			t.Errorf("test query parameter = %q, want a value&more", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"data":{"value":42}}`)
	}))
	defer server.Close()

	httpClient := server.Client()
	httpClient.Timeout = 5 * time.Second
	client := coinmarketcap.NewClient(cmchttp.NewClientWithHTTPClient(server.URL+"/", "test-key", httpClient))
	metrics := client.GlobalMetrics()
	params := map[string]string{"test": "a value&more"}

	checkResult := func(t *testing.T, result map[string]interface{}, err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
		data, ok := result["data"].(map[string]interface{})
		if !ok || data["value"] != float64(42) {
			t.Fatalf("response = %#v, want data.value = 42", result)
		}
	}
	t.Run("with context", func(t *testing.T) {
		result, err := metrics.FearAndGreedLatestWithContext(context.Background(), params)
		checkResult(t, result, err)
	})
	t.Run("original method", func(t *testing.T) {
		result, err := metrics.FearAndGreedLatest(params)
		checkResult(t, result, err)
	})
}

func TestFearAndGreedLatestWithContextCancelsRequest(t *testing.T) {
	t.Parallel()

	started := make(chan struct{})
	canceled := make(chan struct{})
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		close(started)
		select {
		case <-r.Context().Done():
			close(canceled)
		case <-release:
		}
	}))
	defer server.Close()
	defer close(release)

	client := coinmarketcap.NewClient(cmchttp.NewClientWithHTTPClient(server.URL, "test-key", server.Client()))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() {
		_, err := client.GlobalMetrics().FearAndGreedLatestWithContext(ctx, nil)
		done <- err
	}()

	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	select {
	case <-started:
	case err := <-done:
		t.Fatalf("request returned before reaching the server: %v", err)
	case <-timer.C:
		t.Fatal("request did not reach the server")
	}

	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("request error = %v, want context.Canceled", err)
		}
	case <-timer.C:
		t.Fatal("request did not return after cancellation")
	}
	select {
	case <-canceled:
	case <-timer.C:
		t.Fatal("server did not observe request cancellation")
	}
}
