package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Reporter collects and displays results with thread-safe access
type Reporter struct {
	mode            string
	totalJobs       int
	processedCount  int32 // Use atomic operations for thread safety
	successCount    int32
	failCount       int32
	checkResults    []CheckResult
	downloadResults []DownloadResult
	mutex           sync.Mutex // Protects slices
	startTime       time.Time
}

// NewReporter creates a new reporter
func NewReporter(mode string, totalJobs int) *Reporter {
	return &Reporter{
		mode:            mode,
		totalJobs:       totalJobs,
		checkResults:    make([]CheckResult, 0, totalJobs),
		downloadResults: make([]DownloadResult, 0, totalJobs),
		startTime:       time.Now(),
	}
}

// StartProgressMonitor launches a goroutine to display progress
func (r *Reporter) StartProgressMonitor() chan bool {
	done := make(chan bool)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				r.displayProgress()
			case <-done:
				return
			}
		}
	}()

	return done
}

// displayProgress shows current progress
func (r *Reporter) displayProgress() {
	processed := atomic.LoadInt32(&r.processedCount)
	success := atomic.LoadInt32(&r.successCount)
	failed := atomic.LoadInt32(&r.failCount)

	elapsed := time.Since(r.startTime)
	rate := float64(processed) / elapsed.Seconds()

	fmt.Printf("\r[Progress] Processed: %d/%d | Success: %d | Failed: %d | Rate: %.1f/s",
		processed, r.totalJobs, success, failed, rate)
}

// RecordCheckResult records a check result (thread-safe)
func (r *Reporter) RecordCheckResult(result CheckResult) {
	// Update counters atomically
	atomic.AddInt32(&r.processedCount, 1)
	if result.Success {
		atomic.AddInt32(&r.successCount, 1)
	} else {
		atomic.AddInt32(&r.failCount, 1)
	}

	// Store result with mutex protection
	r.mutex.Lock()
	r.checkResults = append(r.checkResults, result)
	r.mutex.Unlock()
}

// RecordDownloadResult records a download result (thread-safe)
func (r *Reporter) RecordDownloadResult(result DownloadResult) {
	// Update counters atomically
	atomic.AddInt32(&r.processedCount, 1)
	if result.Success {
		atomic.AddInt32(&r.successCount, 1)
	} else {
		atomic.AddInt32(&r.failCount, 1)
	}

	// Store result with mutex protection
	r.mutex.Lock()
	r.downloadResults = append(r.downloadResults, result)
	r.mutex.Unlock()
}

// DisplaySummary shows final results
func (r *Reporter) DisplaySummary() {
	fmt.Println("\n\n=== Summary ===")

	elapsed := time.Since(r.startTime)
	fmt.Printf("Total Time: %.2f seconds\n", elapsed.Seconds())
	fmt.Printf("Total URLs: %d\n", r.totalJobs)
	fmt.Printf("Successful: %d\n", atomic.LoadInt32(&r.successCount))
	fmt.Printf("Failed: %d\n", atomic.LoadInt32(&r.failCount))
	fmt.Printf("Success Rate: %.1f%%\n",
		float64(atomic.LoadInt32(&r.successCount))/float64(r.totalJobs)*100)

	if r.mode == "check" {
		r.displayCheckSummary()
	} else {
		r.displayDownloadSummary()
	}
}

// displayCheckSummary shows detailed check results
func (r *Reporter) displayCheckSummary() {
	fmt.Println("\n--- Status Code Distribution ---")

	statusCodes := make(map[int]int)
	r.mutex.Lock()
	for _, result := range r.checkResults {
		if result.Success {
			statusCodes[result.StatusCode]++
		}
	}
	r.mutex.Unlock()

	for code, count := range statusCodes {
		fmt.Printf("  HTTP %d: %d URLs\n", code, count)
	}

	// Show first few failures
	fmt.Println("\n--- Sample Failures ---")
	r.mutex.Lock()
	failCount := 0
	for _, result := range r.checkResults {
		if !result.Success && failCount < 5 {
			fmt.Printf("  [%d] %s - %s\n", result.Job.ID, result.Job.URL, result.Error)
			failCount++
		}
	}
	r.mutex.Unlock()

	if failCount == 0 {
		fmt.Println("  (No failures)")
	} else if int(atomic.LoadInt32(&r.failCount)) > 5 {
		fmt.Printf("  ... and %d more failures\n", int(atomic.LoadInt32(&r.failCount))-5)
	}
}

// displayDownloadSummary shows detailed download results
func (r *Reporter) displayDownloadSummary() {
	fmt.Println("\n--- Download Statistics ---")

	var totalSize int64
	r.mutex.Lock()
	for _, result := range r.downloadResults {
		if result.Success {
			totalSize += int64(result.ContentSize)
		}
	}
	r.mutex.Unlock()

	fmt.Printf("  Total Downloaded: %.2f MB\n", float64(totalSize)/(1024*1024))
	fmt.Printf("  Average Size: %.2f KB\n",
		float64(totalSize)/float64(atomic.LoadInt32(&r.successCount))/1024)

	// Show first few failures
	fmt.Println("\n--- Sample Failures ---")
	r.mutex.Lock()
	failCount := 0
	for _, result := range r.downloadResults {
		if !result.Success && failCount < 5 {
			fmt.Printf("  [%d] %s - %s\n", result.Job.ID, result.Job.URL, result.Error)
			failCount++
		}
	}
	r.mutex.Unlock()

	if failCount == 0 {
		fmt.Println("  (No failures)")
	} else if int(atomic.LoadInt32(&r.failCount)) > 5 {
		fmt.Printf("  ... and %d more failures\n", int(atomic.LoadInt32(&r.failCount))-5)
	}
}
