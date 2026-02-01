package main

import (
	"sync"
)

// Job represents a URL to be processed
type Job struct {
	ID  int
	URL string
}

// CheckResult represents the result of checking a URL's status
type CheckResult struct {
	Job        Job
	StatusCode int
	Success    bool
	Error      string
}

// DownloadResult represents the result of downloading content from a URL
type DownloadResult struct {
	Job         Job
	Content     []byte
	ContentSize int
	Success     bool
	Error       string
}

// WorkerPool manages concurrent execution of jobs
type WorkerPool struct {
	maxWorkers      int
	jobs            chan Job
	checkResults    chan CheckResult
	downloadResults chan DownloadResult
	wg              sync.WaitGroup
	mode            string
	retries         int
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(maxWorkers int, mode string, retries int) *WorkerPool {
	return &WorkerPool{
		maxWorkers:      maxWorkers,
		jobs:            make(chan Job, maxWorkers*2), // Buffered channel for smoother flow
		checkResults:    make(chan CheckResult, maxWorkers*2),
		downloadResults: make(chan DownloadResult, maxWorkers*2),
		mode:            mode,
		retries:         retries,
	}
}

// Start launches worker goroutines
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.maxWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker is the goroutine that processes jobs
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for job := range wp.jobs {
		if wp.mode == "check" {
			result := checkURL(job, wp.retries)
			wp.checkResults <- result
		} else if wp.mode == "download" {
			result := downloadURL(job, wp.retries)
			wp.downloadResults <- result
		}
	}
}

// Submit adds a job to the worker pool
func (wp *WorkerPool) Submit(job Job) {
	wp.jobs <- job
}

// Close signals that no more jobs will be submitted
func (wp *WorkerPool) Close() {
	close(wp.jobs)
}

// Wait blocks until all workers have finished processing
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.checkResults)
	close(wp.downloadResults)
}

// GetCheckResults returns the channel for check results
func (wp *WorkerPool) GetCheckResults() <-chan CheckResult {
	return wp.checkResults
}

// GetDownloadResults returns the channel for download results
func (wp *WorkerPool) GetDownloadResults() <-chan DownloadResult {
	return wp.downloadResults
}
