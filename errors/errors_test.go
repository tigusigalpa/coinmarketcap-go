package errors

import (
	"errors"
	"testing"
)

func TestAPIError(t *testing.T) {
	t.Parallel()

	cause := errors.New("network failure")
	response := map[string]interface{}{"status": "failed"}
	err := NewAPIError("request failed", 502, response, cause)
	if got, want := err.Error(), "API Error [502]: request failed"; got != want {
		t.Fatalf("Error() = %q, want %q", got, want)
	}
	if !errors.Is(err, cause) {
		t.Fatal("Unwrap() did not return the underlying error")
	}
	if err.GetResponse()["status"] != "failed" {
		t.Fatalf("GetResponse() = %#v", err.GetResponse())
	}
}

func TestSpecializedErrors(t *testing.T) {
	t.Parallel()

	response := map[string]interface{}{"status": "failed"}
	retryAfter := 15
	cases := []struct {
		name string
		err  *APIError
		code int
	}{
		{"authentication", NewAuthenticationError("invalid key", response).APIError, 401},
		{"rate limit", NewRateLimitError("slow down", &retryAfter, response).APIError, 429},
		{"invalid request", NewInvalidRequestError("bad input", response).APIError, 400},
		{"not found", NewNotFoundError("missing", response).APIError, 404},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err.StatusCode != tc.code || tc.err.Response["status"] != "failed" {
				t.Fatalf("error = %#v", tc.err)
			}
		})
	}

	if got := NewRateLimitError("slow down", &retryAfter, response).GetRetryAfter(); got == nil || *got != retryAfter {
		t.Fatalf("GetRetryAfter() = %v", got)
	}
}
