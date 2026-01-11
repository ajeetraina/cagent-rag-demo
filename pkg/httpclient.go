package pkg

import (
	"math"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	maxRetries int
	baseDelay  time.Duration
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		maxRetries: 3,
		baseDelay:  100 * time.Millisecond,
	}
}

// Do executes the request with exponential backoff retry
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	var lastErr error
	
	for attempt := 0; attempt < c.maxRetries; attempt++ {
		resp, err := c.httpClient.Do(req)
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}
		
		lastErr = err
		if resp != nil {
			resp.Body.Close()
		}

		// Exponential backoff: 100ms, 200ms, 400ms...
		delay := c.baseDelay * time.Duration(math.Pow(2, float64(attempt)))
		time.Sleep(delay)
	}

	return nil, lastErr
}
