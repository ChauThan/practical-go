package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	// CLI flags
	urlFile := flag.String("file", "urls.txt", "Path to file containing URLs (one per line)")
	maxWorkers := flag.Int("workers", 10, "Maximum number of concurrent workers")
	mode := flag.String("mode", "check", "Operation mode: 'check' (status codes) or 'download' (fetch content)")
	retries := flag.Int("retries", 3, "Number of retry attempts for failed requests")

	flag.Parse()

	// Validate inputs
	if *maxWorkers < 1 {
		fmt.Println("Error: workers must be at least 1")
		os.Exit(1)
	}

	if *mode != "check" && *mode != "download" {
		fmt.Println("Error: mode must be 'check' or 'download'")
		os.Exit(1)
	}

	// Display configuration
	fmt.Println("=== Concurrent Website Checker & Downloader ===")
	fmt.Printf("Configuration:\n")
	fmt.Printf("  URL File: %s\n", *urlFile)
	fmt.Printf("  Max Workers: %d\n", *maxWorkers)
	fmt.Printf("  Mode: %s\n", *mode)
	fmt.Printf("  Retries: %d\n", *retries)
	fmt.Println("==============================================")

	// Read URLs from file
	fmt.Println("\nReading URLs from file...")
	urls, err := ReadURLsFromFile(*urlFile)
	if err != nil {
		fmt.Printf("Error reading URLs: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully loaded %d URLs\n", len(urls))

	// Create worker pool
	pool := NewWorkerPool(*maxWorkers, *mode, *retries)

	// Create reporter
	reporter := NewReporter(*mode, len(urls))

	// Start progress monitoring
	progressDone := reporter.StartProgressMonitor()

	// Start workers
	pool.Start()

	// Submit jobs
	fmt.Println("\nSubmitting jobs to worker pool...")
	for i, url := range urls {
		pool.Submit(Job{ID: i + 1, URL: url})
	}

	// Signal no more jobs
	pool.Close()

	// Collect results in a separate goroutine
	var collectorWG sync.WaitGroup
	collectorWG.Add(1)
	go func() {
		defer collectorWG.Done()
		if *mode == "check" {
			for result := range pool.GetCheckResults() {
				reporter.RecordCheckResult(result)
			}
		} else {
			for result := range pool.GetDownloadResults() {
				reporter.RecordDownloadResult(result)
			}
		}
	}()

	// Wait for all workers to finish
	pool.Wait()

	// Wait for result collector to finish
	collectorWG.Wait()

	// Stop progress monitor
	close(progressDone)

	// Display final summary
	reporter.DisplaySummary()

	fmt.Println("\n=== Processing Complete ===")
}
