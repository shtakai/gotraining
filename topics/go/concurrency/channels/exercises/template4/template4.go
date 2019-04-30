// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Add imports.
// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the channel for sharing results.
	results := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	shutdown := make(chan int)

	// Define the size of the worker pool. Use runtime.NumCPU to size the pool based on number of processors.
	goroutines := runtime.NumCPU()

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(goroutines)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i :=0; i < goroutines; i++{
		go func(){
			//defer wg.Done()

			// Start an infinite loop.
			for true{

				// Generate a random number up to 1000.
				n := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select
				{
					// In one case send the random number.
					case results <- n:
						fmt.Println("number")
						results <- n

					// In another case receive from the shutdown channel.
					case <- shutdown:
						fmt.Println("received shutdown")
						shutdown <- n


				}
			}
		}()
	}

	// Create a slice to hold the random numbers.
	var randNumbers []int


	// Receive from the values channel with range.
	for ch := range results{

		// continue the loop if the value was even.
		if ch % 2 == 0 {
			continue
		}

		// Store the odd number.
		randNumbers = append(randNumbers, ch)

		// break the loop once we have 100 results.
		if len(results) == 100 {
            break
		}
	}

	// Send the shutdown signal by closing the shutdown channel.
	close(shutdown)

	// Wait for the Goroutines to finish.
	wg.Wait()

	// Print the values in our slice.
	fmt.Println("finished")
	fmt.Println("slice:", randNumbers)
}
