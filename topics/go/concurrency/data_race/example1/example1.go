// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build -race

// Sample program to show how to create race conditions in
// our programs. We don't want to do this.
package main

import (
	"fmt"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int

func main() {

	// Number of goroutines to use.
	const grs = 2000

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// go build
	//-rwxr-xr-x  1 alyson  staff  2097712 Apr 30 11:17 example1
	//-rw-r--r--  1 alyson  staff     1913 Apr 30 11:15 example1.go

	// go build --race
	//-rwxr-xr-x  1 alyson  staff  2690284 Apr 30 11:18 example1 (watch)
	//-rw-r--r--  1 alyson  staff     2040 Apr 30 11:18 example1.go

	// Create two goroutines.
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {

				// Capture the value of Counter.
				value := counter // G7 Reading

				// Increment our local value of Counter.
				value++

				// Store the value back into Counter.
				counter = value // G6 Writing
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
// Final Counter: 4
// works my machine


// grs => 2000 にする（production
//Final Counter: 3586
//Final Counter: 3534
//Final Counter: 3560
// ???????

// => data race

// for production w/watching condition? => it depends
// カナリーリリースでも
// race detector pay cost. (watching condition)

// XXXXXXXXXXXXXXXXXXXXXXXXXXXXXx
//go run --race example1.go
//==================
//WARNING: DATA RACE
//Read at 0x0000011d5318 by goroutine 7:
//  main.main.func1()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:41 +0x47
//
//Previous write at 0x0000011d5318 by goroutine 6:
//  main.main.func1()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:47 +0x63
//
//Goroutine 7 (running) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:37 +0xa2
//
//Goroutine 6 (finished) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:37 +0xa2
//==================
//Final Counter: 4000
//Found 1 data race(s)
//exit status 66

// どうなおす？
//  => ex2へ

/*
==================
WARNING: DATA RACE
Read at 0x0000011a5118 by goroutine 7:
  main.main.func1()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:33 +0x4e

Previous write at 0x0000011a5118 by goroutine 6:
  main.main.func1()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:39 +0x6d

Goroutine 7 (running) created at:
  main.main()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:43 +0xc3

Goroutine 6 (finished) created at:
  main.main()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:43 +0xc3
==================
Final Counter: 4
Found 1 data race(s)
*/
