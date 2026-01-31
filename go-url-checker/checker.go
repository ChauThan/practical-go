package main

import (
	"fmt"
	"net/http"
	"time"
)

// checkURL performs an HTTP GET request to check the status code
// Implements retry logic for transient failures
func checkURL(job Job, maxRetries int) CheckResult {
	client := &http.Client{
		Timeout: 10 * time.Second, // 10 second timeout
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Follow up to 5 redirects
			if len(via) >= 5 {
				return fmt.Errorf("too many redirects")
			}
			return nil
		},
	}

	var lastErr error
	var statusCode int

	// Retry loop with exponential backoff
	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff: 1s, 2s, 4s, ...
			backoff := time.Duration(1<<uint(attempt-1)) * time.Second
			time.Sleep(backoff)
		}

		resp, err := client.Get(job.URL)
		if err != nil {
			lastErr = err
			continue // Retry on error
		}

		statusCode = resp.StatusCode
		resp.Body.Close() // Important: close response body

		// Success - return result
		return CheckResult{
			Job:        job,
			StatusCode: statusCode,
			Success:    true,
			Error:      "",
		}
	}

	// All retries failed
	return CheckResult{
		Job:        job,
		StatusCode: 0,
		Success:    false,
		Error:      fmt.Sprintf("failed after %d retries: %v", maxRetries+1, lastErr),
	}
}
