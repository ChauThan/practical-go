package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int, name string){
	fmt.Printf("Worker %d: Making %s...\n", id, name)

	// simulate work task take 3 seconds
	// but can be cancelled via context
	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("Worker %d: %s is done!\n", id, name)
	case <-ctx.Done():
		fmt.Printf("Worker %d: %s was cancelled!\n", id, name)
	}
} 

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// always call cancel to release resources
	defer cancel()

	go worker(ctx, 1, "Coffee")

	time.Sleep(4 * time.Second) // wait before starting next worker
	fmt.Println("Main: Ending program, cancelling context...")
}