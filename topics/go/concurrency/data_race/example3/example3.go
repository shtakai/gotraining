// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use a mutex to define critical
// sections of code that need synchronous access.
package main

import (
	"fmt"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int // shared memory

// mutex is used to define a critical section of code.
// mutal exclusion lock
//  put some critical section
var mutex sync.Mutex

func main() {

	// Number of goroutines to use.
	const grs = 20000

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				// CRITICAL SECTION start

				// Only allow one goroutine through this critical section at a time.
				mutex.Lock() // LOCK
				// 他のgoroutineはブロックされる
				// accessing shared memory
				{ // <= rubyでいうところのblock
					// Capture the value of counter.
					value := counter

					// Increment our local value of counter.
					value++

					// Store the value back into counter.
					counter = value
				}
				// ↑ critical section w/ {, }

				// UNLOCKしないといけない
				mutex.Unlock()
				// Release the lock and allow any waiting goroutine through.

				// CRITICAL SECTION end
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

// 200
//go run --race example3.go
//Final Counter: 400

// 20000
//go run --race example3.go                                      1 ↵
//Final Counter: 40000

// 時間はかかる