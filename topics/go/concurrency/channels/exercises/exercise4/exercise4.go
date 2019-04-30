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
)

func main() {

	// Create the channel for sharing results.
	values := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	shutdown := make(chan struct{})
	// what type should we apply?
	// bool? int????   no need but need something
	//   => struct{} (empty struct / no fields, 0 byte)

	// Define the size of the worker pool. Use runtime.NumCPU to size the pool based on number of processors.
	poolSize := runtime.NumCPU()

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(poolSize)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i := 0; i < poolSize; i++ {
		go func(id int) { //create goroutine

			// Start an infinite loop.
			for {

				// Generate a random number up to 1000.
				n := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select {

				// In one case send the random number.
				case values <- n: // trying to send to channel
					fmt.Printf("Worker %d sent %d\n", id, n)

					// OR
				// In another case receive from the shutdown channel.
				case <-shutdown: // time to go home
					fmt.Printf("Worker %d shutting down\n", id)
					wg.Done() // tell the WG
					return // leave loop
				}
			}
		}(i)
	}

	// Create a slice to hold the random numbers.
	var nums []int
	for i := range values {

		// continue the loop if the value was even.
		if i%2 == 0 { // don't like
			fmt.Println("Discarding", i)
			continue // go back to loop top
		}

		// Store the odd number.
		fmt.Println("Keeping", i)
		nums = append(nums, i)

		// break the loop once we have 100 results.
		if len(nums) == 100 {
			break
		}
	}

	// Send the shutdown signal by closing the channel.
	fmt.Println("Receiver sending shutdown signal")
	close(shutdown) // send shutdown

	// Wait for the Goroutines to finish.
	wg.Wait() // wait

	// Print the values in our slice.
	fmt.Printf("Result count: %d\n", len(nums))
	fmt.Println(nums)
}
//Worker 7 sent 887
//Keeping 887
//Keeping 81
//Keeping 847
//Keeping 81
//Keeping 59
//Discarding 540
//Worker 1 sent 81
// :
// :
// :
//Discarding 578
//Discarding 154
//Worker 0 sent 10
//Worker 5 sent 720
//Worker 7 sent 700
//Worker 6 sent 565
//Receiver sending shutdown signal => close(shutdown)
//Worker 6 shutting down
//Worker 4 sent 870
//Worker 4 shutting down
//Worker 1 sent 305
//Worker 1 shutting down
//Worker 5 shutting down
//Worker 7 shutting down
//Worker 3 sent 984
//Worker 3 shutting down
//Worker 0 shutting down
//Worker 2 sent 359
//Worker 2 shutting down
//Result count: 100
//[887 81 847 81 59 425 511 89 211 445 237 495 47 947 287 15 541 387 831 429 737 631 485 413 563 433 147 159 353 957 721 189 199 705 703 355 451 605 561 783 563 447 577 463 623 953 137 133 241 59 33 643 891 107 503 843 205 425 351 515 757 687 285 553 591 297 267 137 271 981 79 493 819 981 175 885 387 749 903 547 839 51 351 305 183 801 767 231 223 743 535 657 415 371 39 513 359 783 247 565]
//(