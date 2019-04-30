// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

// Add imports.

// others can communicate via ch.
// using channel, one access the shared memory.

func main() {

	// Create an unbuffered channel.
	ch := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch the goroutine and handle Done.
	go func() {
		goroutine()
		wg.Done()
	}()

	// Launch the goroutine and handle Done.
	go func() {
		goroutine()
		wg.Done()
	}()

	// Send a value to start the counting.
	fmt.Println("send channel 1")
	ch <- 1
	fmt.Println("sent channel 1")

	// Wait for the program to finish.
	fmt.Println("wait")
	wg.Wait()
	fmt.Println("wait->finished")

}

// goroutine simulates sharing a value.
func goroutine( /* parameters */ ) {
	for {

		// Wait for the value to be sent.
		fmt.Println("wait for the value to be sent")
		// If the channel was closed, return.

		if  {//channel is closed

			fmt.Println("if the channel was closed, return")
		}

		// Display the value
		fmt.Println("display the value", value)
		// Terminate when the value is 10.
		if value == 10 {
			fmt.Println("terminate when the value is 10")
            return
		}
		// Increment the value and send it
		// over the channel.
		fmt.Println("increment the value and send it")
		ch <- (value + 1)
	}
}
