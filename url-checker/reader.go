package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// ReadURLsFromFile reads URLs from a text file and validates them
// Returns a slice of valid URLs and any error encountered
func ReadURLsFromFile(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	lineNum := 0

	// Read file line by line
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Validate URL format
		if err := validateURL(line); err != nil {
			fmt.Printf("Warning: Skipping invalid URL at line %d: %s (%v)\n", lineNum, line, err)
			continue
		}

		urls = append(urls, line)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if len(urls) == 0 {
		return nil, fmt.Errorf("no valid URLs found in file")
	}

	return urls, nil
}

// validateURL checks if the URL has valid HTTP/HTTPS scheme
func validateURL(rawURL string) error {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("parse error: %w", err)
	}

	// Check for valid scheme
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("invalid scheme '%s', must be http or https", parsedURL.Scheme)
	}

	// Check for host
	if parsedURL.Host == "" {
		return fmt.Errorf("missing host")
	}

	return nil
}
