# Concurrent Website Checker & Downloader

A Go application that demonstrates concurrent programming patterns by checking HTTP status codes and downloading content from multiple URLs simultaneously.

## Features

- **Concurrent Processing**: Uses goroutines to process multiple URLs in parallel
- **Bounded Concurrency**: Worker pool pattern limits simultaneous operations
- **Two Modes**:
  - **Check Mode**: Verifies HTTP status codes (200, 404, etc.)
  - **Download Mode**: Downloads and stores content in memory
- **Retry Logic**: Automatic retries with exponential backoff for transient failures
- **Real-time Progress**: Live updates showing processing status
- **Thread-Safe**: Uses channels for communication and mutex for shared state

## Go Concurrency Patterns Demonstrated

This project showcases key Go concurrency primitives:

- **Goroutines**: Lightweight threads for parallel execution (`go worker()`)
- **Channels**: Type-safe communication between goroutines
- **WaitGroups**: Synchronization to wait for goroutine completion
- **Mutex**: Thread-safe access to shared result aggregation
- **Worker Pool**: Limits concurrent operations to prevent resource exhaustion

## Installation

```bash
# Clone the repository
cd url-checker

# Build the application
go build -o url-checker

# Or run directly
go run .
```

## Usage

### Basic Usage

```bash
# Check URLs with default settings (10 workers)
go run . -file urls.txt

# Specify number of workers
go run . -file urls.txt -workers 20

# Download mode
go run . -file urls.txt -mode download

# Custom retry attempts
go run . -file urls.txt -retries 5
```

### Command-Line Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-file` | `urls.txt` | Path to file containing URLs (one per line) |
| `-workers` | `10` | Maximum number of concurrent workers |
| `-mode` | `check` | Operation mode: `check` or `download` |
| `-retries` | `3` | Number of retry attempts for failed requests |

### Input File Format

Create a text file with one URL per line:

```
https://www.google.com
https://www.github.com
https://www.stackoverflow.com
# This is a comment (lines starting with # are ignored)

https://www.wikipedia.org
```

## Examples

### Example 1: Quick Status Check

```bash
go run . -file test-urls.txt -workers 5
```

Output:
```
=== Concurrent Website Checker & Downloader ===
Configuration:
  URL File: test-urls.txt
  Max Workers: 5
  Mode: check
  Retries: 3
==============================================

Successfully loaded 7 URLs

[Progress] Processed: 7/7 | Success: 5 | Failed: 2 | Rate: 0.5/s

=== Summary ===
Total Time: 14.23 seconds
Total URLs: 7
Successful: 5
Failed: 2
Success Rate: 71.4%

--- Status Code Distribution ---
  HTTP 200: 3 URLs
  HTTP 403: 2 URLs
```

### Example 2: Download Content

```bash
go run . -file test-urls.txt -mode download -workers 3
```

Output:
```
=== Summary ===
Total Time: 15.42 seconds
Successful: 3
Failed: 4

--- Download Statistics ---
  Total Downloaded: 1.06 MB
  Average Size: 363.13 KB
```

## Project Structure

```
url-checker/
├── main.go          # Entry point and orchestration
├── reader.go        # URL file reading and validation
├── worker.go        # Worker pool implementation
├── checker.go       # HTTP status checking logic
├── downloader.go    # Content downloading logic
├── reporter.go      # Progress reporting and results
├── urls.txt         # Sample URL list
└── test-urls.txt    # Small test URL list
```

## Technical Details

### Concurrency Model

The application uses a **worker pool pattern**:

1. **Job Submission**: Main goroutine submits jobs to a buffered channel
2. **Worker Goroutines**: N workers pull jobs from the channel concurrently
3. **Result Collection**: Separate goroutine collects results from result channels
4. **Synchronization**: WaitGroups ensure all workers complete before exit

```
                    ┌─────────┐
    URLs ────────> │  Jobs   │ ────> Worker 1 ──┐
                   │ Channel │ ────> Worker 2 ──┤
                   └─────────┘ ────> Worker 3 ──┼──> Results
                                    ...          │
                               ────> Worker N ──┘
```

### Performance Characteristics

- **Memory**: ~2KB per goroutine (vs ~1MB for OS threads)
- **Scalability**: Can easily handle 100+ concurrent workers
- **Throughput**: Processes 50-100 URLs per minute (network-dependent)
- **Download Limit**: 10MB per file to prevent memory exhaustion

### Error Handling

- **Network Errors**: Automatic retry with exponential backoff (1s, 2s, 4s...)
- **Timeouts**: 10s for status checks, 30s for downloads
- **HTTP Errors**: Non-2xx status codes trigger retries
- **Invalid URLs**: Skipped during file reading with warnings

## Learning Resources

This project demonstrates concepts covered in:

- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Context and Cancellation](https://go.dev/blog/context)