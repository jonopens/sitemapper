package http

import (
	"fmt"
	"net/http"
	"time"
)

// RetryClient is an HTTP client with retry logic
type RetryClient struct {
	MaxRetries int
	Timeout    time.Duration
	client     *http.Client
}

// NewRetryClient creates a new retry client
func NewRetryClient(maxRetries int, timeout time.Duration) *RetryClient {
	return &RetryClient{
		MaxRetries: maxRetries,
		Timeout:    timeout,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get performs a GET request with retry logic
func (c *RetryClient) Get(url string) (*http.Response, error) {
	var lastErr error
	
	for i := 0; i < c.MaxRetries; i++ {
		resp, err := c.client.Get(url)
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}
		
		if err != nil {
			lastErr = err
		} else {
			lastErr = fmt.Errorf("server error: %d", resp.StatusCode)
			resp.Body.Close()
		}
		
		// Exponential backoff
		if i < c.MaxRetries-1 {
			backoff := time.Duration(i+1) * time.Second
			time.Sleep(backoff)
		}
	}
	
	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

// Post performs a POST request with retry logic
func (c *RetryClient) Post(url, contentType string, body interface{}) (*http.Response, error) {
	// TODO: Implement POST with retry logic
	return nil, fmt.Errorf("not implemented")
}

