package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Report when the worker is done

	for job := range jobs {
		fmt.Printf("Worker %d: Making %s...\n", id, job)
		time.Sleep(500 * time.Millisecond) // Simulate time taken to make the drink 0.5 seconds
		results <- fmt.Sprintf("%s is ready!", job)
	}
}

func main() {
	start := time.Now()

	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan string, numWorkers * 2)
	results := make(chan string, numWorkers * 2)
	var wg sync.WaitGroup

	// Start 3 workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1);
		go worker(w, jobs, results, &wg)
	}

	// Send 10 drink orders
	drinks := []string{"Coffee", "Tea", "Juice", "Soda", "Water", "Milk", "Smoothie", "Lemonade", "Hot Chocolate", "Espresso"}
	for _, drink := range drinks {
		jobs <- drink
	}
	close(jobs) // Close the jobs channel since no more jobs will be sent

	go func() {
		wg.Wait()
		close(results) // Close results channel when all workers are done
	}()

	// Collect results
	for result := range results {
		fmt.Println(result)
	}

	elapsed := time.Since(start)
	fmt.Printf("All drinks are ready in %s\n", elapsed)
}