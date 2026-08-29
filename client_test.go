package coinmarketcap

import (
	"sync"
	"testing"
	"time"
)

func TestClientBuilder(t *testing.T) {
	t.Parallel()

	_, err := NewClientBuilder().Build()
	if err == nil {
		t.Fatal("Build() error = nil, want missing API key error")
	}

	client, err := NewClientBuilder().SetAPIKey("test-key").SetTimeout(time.Second).UseSandbox().Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	cryptocurrency := client.Cryptocurrency()
	exchange := client.Exchange()
	globalMetrics := client.GlobalMetrics()
	tools := client.Tools()
	if client.Cryptocurrency() != cryptocurrency || client.Exchange() != exchange || client.GlobalMetrics() != globalMetrics || client.Tools() != tools {
		t.Fatal("service accessors must cache their service")
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = client.Cryptocurrency()
			_ = client.Exchange()
			_ = client.GlobalMetrics()
			_ = client.Tools()
		}()
	}
	wg.Wait()
}
