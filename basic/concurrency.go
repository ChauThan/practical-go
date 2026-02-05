package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	done := make(chan string)

	go makeDrink("Coffee", 2*time.Second, done)
	go makeDrink("Tea", 1*time.Second, done)

	for i := 0; i < 2; i++ {
		fmt.Println(<-done)
	}

	elapsed := time.Since(start)
	fmt.Printf("All drinks are ready in %s\n", elapsed)
}

func makeDrink(name string, duration time.Duration, done chan<- string) {
	fmt.Printf("Making %s...\n", name)

	time.Sleep(duration)
	done <- fmt.Sprintf("%s is ready!", name)
}