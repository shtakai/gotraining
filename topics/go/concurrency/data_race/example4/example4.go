// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use a read/write mutex to define critical
// sections of code that needs synchronous access.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// data is a slice that will be shared.
var data []string // shared memory

// rwMutex is used to define a critical section of code.
var rwMutex sync.RWMutex // read write mutex

// Number of reads occurring at ay given time.
var readCount int64 // keep track goroutines#

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(1)

	// Create a writer goroutine.
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writer(i)
		}
		wg.Done()
	}()

	// Create eight reader goroutines.
	// here have very common bug
	for i := 0; i < 8; i++ { // only one variables: i
		i := i // これでもOK。 左と右のiは違う
		go func() {
			// infinite loop
			for {
				reader(i)
			}
		}()
	}

	//for i := 0; i < 8; i++ { // only one variables: i
	//	go func(i int) { // functionパラメータを 明示してもOK
	//		// infinite loop
	//		for {
	//			reader(i)
	//		}
	//	}(i)
	//}
	// go run -race example4.go >/dev/null

	// Wait for the write goroutine to finish.
	wg.Wait()
	fmt.Println("Program Complete")
}

// writer adds a new string to the slice in random intervals.
func writer(i int) {

	// Only allow one goroutine to read/write to the slice at a time.
	rwMutex.Lock()
	{
		// Capture the current read count.
		// Keep this safe though we can due without this call.
		rc := atomic.LoadInt64(&readCount)

		// Perform some work since we have a full lock.
		fmt.Printf("****> : Performing Write : RCount[%d]\n", rc)
		data = append(data, fmt.Sprintf("String: %d", i)) // access shared memory
	}
	rwMutex.Unlock()
	// Release the lock.
}

// reader wakes up and iterates over the data slice.
func reader(id int) {

	// Any goroutine can read when no write operation is taking place.
	rwMutex.RLock() // promise read and not write
	{
		// Increment the read count value by 1.
		rc := atomic.AddInt64(&readCount, 1)

		// Perform some read work and display values.
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("%d : Performing Read : Length[%d] RCount[%d]\n", id, len(data), rc) // access shared memory

		// Decrement the read count value by 1.
		atomic.AddInt64(&readCount, -1)
	}
	rwMutex.RUnlock()
	// Release the read lock.
}

// go run ex
//8 : Performing Read : Length[8] RCount[8]
//8 : Performing Read : Length[8] RCount[8]
//8 : Performing Read : Length[8] RCount[8]
//8 : Performing Read : Length[8] RCount[8]
//****> : Performing Write : RCount[0]
//8 : Performing Read : Length[9] RCount[2]
//8 : Performing Read : Length[9] RCount[5]
//8 : Performing Read : Length[9] RCount[8]
//8 : Performing Read : Length[9] RCount[1]
//8 : Performing Read : Length[9] RCount[7]
//8 : Performing Read : Length[9] RCount[8]
//8 : Performing Read : Length[9] RCount[4]
//8 : Performing Read : Length[9] RCount[6]
//8 : Performing Read : Length[9] RCount[8]
//8 : Performing Read : Length[9] RCount[3]
//8 : Performing Read : Length[9] RCount[8]
//8 : Performing Read : Length[9] RCount[3]
//8 : Performing Read : Length[9] RCount[8]
//****> : Performing Write : RCount[0]
//Program Complete

// go run -race example4.go >/dev/null                            1 ↵
//==================
//WARNING: DATA RACE
//Read at 0x00c42001a0e8 by goroutine 7:
//  main.main.func2()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example4/example4.go:50 +0x38
//
//Previous write at 0x00c42001a0e8 by main goroutine:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example4/example4.go:46 +0x10b
//
//Goroutine 7 (running) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example4/example4.go:47 +0xe7
//==================
//Found 1 data race(s)
//exit status 66
//(3.6.5/envs/3.6.5) alyson@alysonX1

//» go run example4.go |head
//8 : Performing Read : Length[0] RCount[6]
//8 : Performing Read : Length[0] RCount[8]
//3 : Performing Read : Length[0] RCount[1]
//8 : Performing Read : Length[0] RCount[7]
//8 : Performing Read : Length[0] RCount[8]
//8 : Performing Read : Length[0] RCount[8]
//8 : Performing Read : Length[0] RCount[4]
//8 : Performing Read : Length[0] RCount[2]
//8 : Performing Read : Length[0] RCount[3]
//8 : Performing Read : Length[0] RCount[5]
//signal: broken pipe