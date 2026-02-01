package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const maxDownloadSize = 10 * 1024 * 1024 // 10 MB limit

// downloadURL downloads content from a URL
// Implements retry logic and size limits to prevent memory issues
func downloadURL(job Job, maxRetries int) DownloadResult {
	client := &http.Client{
		Timeout: 30 * time.Second, // 30 second timeout for downloads
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return fmt.Errorf("too many redirects")
			}
			return nil
		},
	}

	var lastErr error

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
		defer resp.Body.Close()

		// Check if status code indicates success
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			lastErr = fmt.Errorf("HTTP %d", resp.StatusCode)
			continue // Retry on non-2xx status
		}

		// Check content length if available
		if resp.ContentLength > maxDownloadSize {
			return DownloadResult{
				Job:         job,
				Content:     nil,
				ContentSize: int(resp.ContentLength),
				Success:     false,
				Error:       fmt.Sprintf("content too large: %d bytes (max: %d)", resp.ContentLength, maxDownloadSize),
			}
		}

		// Read content with size limit
		limitedReader := io.LimitReader(resp.Body, maxDownloadSize+1)
		content, err := io.ReadAll(limitedReader)
		if err != nil {
			lastErr = err
			continue // Retry on read error
		}

		// Check if we exceeded the limit
		if len(content) > maxDownloadSize {
			return DownloadResult{
				Job:         job,
				Content:     nil,
				ContentSize: len(content),
				Success:     false,
				Error:       fmt.Sprintf("content too large: exceeded %d bytes", maxDownloadSize),
			}
		}

		// Success - return result
		return DownloadResult{
			Job:         job,
			Content:     content,
			ContentSize: len(content),
			Success:     true,
			Error:       "",
		}
	}

	// All retries failed
	return DownloadResult{
		Job:         job,
		Content:     nil,
		ContentSize: 0,
		Success:     false,
		Error:       fmt.Sprintf("failed after %d retries: %v", maxRetries+1, lastErr),
	}
}
