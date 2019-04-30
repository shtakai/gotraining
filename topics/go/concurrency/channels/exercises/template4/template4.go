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
	shutdown := make(chan struct{}) // don't care type.  struct{} => empty/0 bytes

	// Define the size of the worker pool. Use runtime.NumCPU to size the pool based on number of processors.
	goroutines := runtime.NumCPU()

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(goroutines)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i :=0; i < goroutines; i++{
		go func(id int){ // use id for recognize as worker id
			//defer wg.Done() <= no need here.  Done calls when shutdown

			// Start an infinite loop.
			for {

				// Generate a random number up to 1000.
				n := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select
				{
				// In one case send the random number.
				case results <- n: // when send n to results
					fmt.Println(id, " send number:", n)


				// In another case receive from the shutdown channel.
				case <- shutdown: // when receive shutdown  | close(shutdown)
					fmt.Println(id, " received shutdown")
					wg.Done() // finish this routine. guaranteed because of shutdown
					return // leave
				}
			}
		}(i) // argument is i => id
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
		if len(randNumbers) == 100 {
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
//7  send number: 483
//5  send number: 129
//5  send number: 211
// :
// :
// :
//3  send number: 127
//3  send number: 735
//1  send number: 194
//7  send number: 175
//7  received shutdown
//2  send number: 14
//5  send number: 646
//5  received shutdown
//6  send number: 627
//6  received shutdown
//3  received shutdown
//2  received shutdown
//4  send number: 324
//4  received shutdown
//1  received shutdown
//0  send number: 719
//0  received shutdown
//finished
//slice: [483 129 195 211 431 185 629 823 787 393 619 571 39 231 881 61 859 773 515 389 629 713 425 677 669 203 27 83 911 927 751 639 997 375 227 525 777 763 271 69 815 361 255 249 393 845 831 763 707 149 989 67 245 199 579 503 253 283 179 987 719 951 411 441 45 785 79 925 29 675 705 877 645 981 343 333 587 589 701 835 49 515 987 883 105 379 161 633 501 719 125 969 349 283 529 159 627 175 127 735]
