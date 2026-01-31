# Quick Reference Guide

## Common Commands

```bash
# Basic check (default 10 workers)
go run .

# Fast check with more workers
go run . -workers 20

# Download mode
go run . -mode download

# Custom file
go run . -file my-urls.txt

# More aggressive retries
go run . -retries 5

# Combined options
go run . -file urls.txt -workers 20 -mode check -retries 3
```

## Creating Your Own URL List

Create a text file (e.g., `my-urls.txt`):

```
# Comments start with #
https://example.com
https://another-site.org

# Empty lines are ignored

https://third-site.net
```

Run with:
```bash
go run . -file my-urls.txt
```

## Understanding the Output

### Progress Line
```
[Progress] Processed: 36/38 | Success: 35 | Failed: 1 | Rate: 4.5/s
```
- **Processed**: Completed / Total URLs
- **Success**: URLs that responded successfully
- **Failed**: URLs that failed after retries
- **Rate**: Processing speed (URLs per second)

### Summary Section
```
=== Summary ===
Total Time: 8.07 seconds
Total URLs: 38
Successful: 35
Failed: 3
Success Rate: 92.1%
```

### Status Code Distribution (Check Mode)
Shows HTTP status codes returned:
```
--- Status Code Distribution ---
  HTTP 200: 31 URLs  (OK)
  HTTP 403: 4 URLs   (Forbidden)
  HTTP 404: 2 URLs   (Not Found)
```

### Download Statistics (Download Mode)
Shows download metrics:
```
--- Download Statistics ---
  Total Downloaded: 1.06 MB
  Average Size: 363.13 KB
```

## Troubleshooting

### High Failure Rate
- **Too many workers**: Reduce with `-workers 5`
- **Network issues**: Check internet connection
- **Rate limiting**: Add delays or reduce workers

### Slow Performance
- **Increase workers**: Try `-workers 20` or higher
- **Reduce retries**: Use `-retries 1` for faster feedback
- **Network latency**: Some sites are naturally slow

### Memory Issues (Download Mode)
- **Reduce workers**: Use `-workers 5`
- **Large files**: Application limits downloads to 10MB each
- **Many URLs**: Process in smaller batches

## Go Concurrency Concepts

### Goroutines
Lightweight threads that run functions concurrently:
```go
go myFunction()  // Launches myFunction in a new goroutine
```

### Channels
Type-safe queues for communication between goroutines:
```go
ch := make(chan int)  // Create channel
ch <- 42              // Send value
value := <-ch         // Receive value
```

### WaitGroups
Synchronization to wait for multiple goroutines:
```go
var wg sync.WaitGroup
wg.Add(1)       // Increment counter
go func() {
    defer wg.Done()  // Decrement when done
    // ... work ...
}()
wg.Wait()       // Block until counter reaches 0
```

### Mutex
Lock for protecting shared data:
```go
var mu sync.Mutex
mu.Lock()
// ... access shared data ...
mu.Unlock()
```

## Performance Tuning

### Worker Count Guidelines
- **Fast network**: 20-50 workers
- **Slow network**: 5-10 workers
- **Local testing**: 3-5 workers
- **Many timeouts**: Reduce workers

### Retry Strategy
- **Stable sites**: `-retries 1`
- **Unreliable sites**: `-retries 5`
- **Production**: `-retries 3` (default)

## Next Steps

1. Try both modes: `-mode check` and `-mode download`
2. Experiment with different worker counts
3. Create custom URL lists for your use case
4. Modify the code to learn Go concurrency
5. Add new features (see README for ideas)

## Related Resources

- [Go by Example: Goroutines](https://gobyexample.com/goroutines)
- [Go by Example: Channels](https://gobyexample.com/channels)
- [Go by Example: Worker Pools](https://gobyexample.com/worker-pools)
- [Effective Go](https://go.dev/doc/effective_go)
