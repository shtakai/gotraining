// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffer channel so no send every blocks. Don't allocate more buffers than
// you need. Have the main goroutine display each random number is receives and
// then terminate the program.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the buffer channel with a buffer for
	// each goroutine to be created.
	values := make(chan int, goroutines) //buffer

	// Iterate and launch each goroutine.
	for gr := 0; gr < goroutines; gr++ {

		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			values <- rand.Intn(1000) // send random to values
			// send to channel == goroutine is finished (no need wg.Done())
		}()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	wait := goroutines //initialize wait

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	var nums []int
	for wait > 0 {
		nums = append(nums, <-values)
		wait--
	}
	//for i :=0 ; i<goroutines; i++{
	//	nums = append(nums, <-values)
	//}

	// Print the values in our slice.
	fmt.Println(nums)
}

//[170 786 796 828 235 83 464 341 100 366 365 561 238 898 32 974 879 442 435 183 180 81 156 782 772 800 68 787 53 532 84 891 15 928 762 610 947 795 264 780 851 442 755 267 433 20 448 255 237 616 763 274 131 864 623 736 398 350 428 737 934 984 637 217 915 65 33 323 93 646 555 16 518 318 697 623 44 711 629 818 492 856 243 161 765 25 843 550 903 394 286 996 269 252 602 314 201 807 71 45]
