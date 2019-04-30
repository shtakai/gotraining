// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses goroutines to generate up to 100 random numbers.
// Do not send values that are divisible by 2. Have the main goroutine receive
// values and add them to a slice.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Declare constant for number of goroutines.
const goroutines = 100

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the channel for sharing results.
	results := make(chan int)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(goroutines)

	// Iterate and launch each goroutine.
	for i := 0; i < goroutines; i++ {

		// Create an anonymous function for each goroutine.
		go func(){

			// Ensure the waitgroup is decremented when this function returns.
			defer wg.Done() // guarantees Done

			// Generate a random number up to 1000.
			number := rand.Intn(100)

			// Return early if the number is even. (n%2 == 0)
            if number % 2 == 0 {
				fmt.Print("e")
				return
			}

			// Send the odd values through the channel.
			fmt.Print("O")
			results <- number
		}()
	}

	fmt.Println("")
	// Create a goroutine that waits for the other goroutines to finish then
	// closes the channel.
	go func() {
		fmt.Println("start waiting")
		wg.Wait() // wait all goroutines Done
		fmt.Println("\nall done")
		fmt.Println("close the channel")
		close(results) // send result that 'no more data'
		fmt.Println("finish close the channel")
	}()
	// ↓大切
	//go func() {
	//	wg.Wait() // wait all goroutines Done
	//	close(results) // send result that 'no more data'
	//}()


	ints := []int{}
	// Receive from the channel until it is closed.
	// Store values in a slice of ints.
	for data := range results {
		fmt.Print(".")
		ints = append(ints, data)
	}

	// Print the values in our slice.
	fmt.Println("LAST: display []int")
	fmt.Println("ints:", ints)
}
//OOOeOOeOOeOOOeOeeeOeOOOeOeOeOOeOeOOOOOOeeeeeOeeeeOOeeeeOeOeOOeeOOeOeOOOOOOeeeOeO
//OO.OeOOeOO.start waiting
//OeeOeO........ee........ee............e........................
//all done
//close the channel
//finish close the channel
//LAST: display []int
//ints: [49 69 83 65 63 31 43 67 89 75 43 33 93 51 87 59 53 11 77 41 29 87 51 65 19 57 49 31 93 69 19 85 3 37 91 87 19 35 29 39 69 15 59 73 73 7 7 9 47 59 97 81 49 45]
