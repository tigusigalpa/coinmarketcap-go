package errors

import "fmt"

// APIError represents an error returned while calling the CoinMarketCap API.
type APIError struct {
	Message    string
	StatusCode int
	Response   map[string]interface{}
	Err        error
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error [%d]: %s", e.StatusCode, e.Message)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

// GetResponse returns the parsed API error response, when available.
func (e *APIError) GetResponse() map[string]interface{} {
	return e.Response
}

// AuthenticationError indicates that the API key is missing or invalid.
type AuthenticationError struct {
	*APIError
}

// NewAuthenticationError creates an authentication error.
func NewAuthenticationError(message string, response map[string]interface{}) *AuthenticationError {
	return &AuthenticationError{
		APIError: &APIError{
			Message:    message,
			StatusCode: 401,
			Response:   response,
		},
	}
}

// RateLimitError indicates that the API request limit has been exceeded.
type RateLimitError struct {
	*APIError
	RetryAfter *int
}

// NewRateLimitError creates a rate limit error.
func NewRateLimitError(message string, retryAfter *int, response map[string]interface{}) *RateLimitError {
	return &RateLimitError{
		APIError: &APIError{
			Message:    message,
			StatusCode: 429,
			Response:   response,
		},
		RetryAfter: retryAfter,
	}
}

// GetRetryAfter returns the number of seconds to wait before retrying, when known.
func (e *RateLimitError) GetRetryAfter() *int {
	return e.RetryAfter
}

// InvalidRequestError indicates invalid request parameters.
type InvalidRequestError struct {
	*APIError
}

// NewInvalidRequestError creates an invalid request error.
func NewInvalidRequestError(message string, response map[string]interface{}) *InvalidRequestError {
	return &InvalidRequestError{
		APIError: &APIError{
			Message:    message,
			StatusCode: 400,
			Response:   response,
		},
	}
}

// NotFoundError indicates that the requested API resource was not found.
type NotFoundError struct {
	*APIError
}

// NewNotFoundError creates a not found error.
func NewNotFoundError(message string, response map[string]interface{}) *NotFoundError {
	return &NotFoundError{
		APIError: &APIError{
			Message:    message,
			StatusCode: 404,
			Response:   response,
		},
	}
}

// NewAPIError creates a general API error.
func NewAPIError(message string, statusCode int, response map[string]interface{}, err error) *APIError {
	return &APIError{
		Message:    message,
		StatusCode: statusCode,
		Response:   response,
		Err:        err,
	}
}
