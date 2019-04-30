// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Fix the race condition in this program.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// numbers maintains a set of random numbers.
var(
	numbers []int // shared
	mutex sync.Mutex // add mutex for data race
	rwMutex sync.RWMutex
)

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Number of goroutines to use.
	const grs = 3

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create three goroutines to generate random numbers.
	for i := 0; i < grs; i++ {
		go func() {
			random(10)
			wg.Done()
		}()
		// 別解
		//go func() {
		//	newslice := random(10)
		//	rwMutex.Lock()
		//	{
		//		numbers = append(numbers, newslice...) // accessing shared memory
		//
		//	}
		//	rwMutex.Unlock()
		//	wg.Done()
		//}()

	}


	// Wait for all the goroutines to finish.
	wg.Wait()

	// Display the set of random numbers.
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}

// random generates random numbers and stores them into a slice.
func random(amount int) {

	// Generate as many random numbers as specified.
	//ここでブロックを入れるとパフォーマンスが悪くなる
	// mutex.Lock() X
	// defer.Unlock() X
	for i := 0; i < amount; i++ {

		n := rand.Intn(100) // ここをブロックに入れるとパフォーマンスが悪くなる
		// here is CRITICAL SECTION
		mutex.Lock()
		{
			// n := rand.Intn(100) XX パフォーマンスが悪くなる
			numbers = append(numbers, n) // access shared memory R/W
		}
		mutex.Unlock()
	}
}

// 別解
//func random(amount int) []int {
//	arrInt := []int{}
//	// generate as many random numbers as specified.
//	for i:=0; i< amount; i++ {
//		n:= rand.Intn(100)
//		arrInt = append(arrInt, n)
//	}
//	return arrInt
//}

//go run -race template1.go > /dev/null               1 ↵
//==================
//WARNING: DATA RACE
//Read at 0x0000011be1b0 by goroutine 7:
//  main.random()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:54 +0x78
//  main.main.func1()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:34 +0x37
//
//Previous write at 0x0000011be1b0 by goroutine 6:
//  main.random()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:54 +0xee
//  main.main.func1()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:34 +0x37
//
//Goroutine 7 (running) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:33 +0xb3
//
//Goroutine 6 (finished) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:33 +0xb3
//==================
//==================
//WARNING: DATA RACE
//Read at 0x00c4200b8048 by goroutine 7:
//  runtime.growslice()
//      /Users/alyson/.goenv/versions/1.10.3/src/runtime/slice.go:89 +0x0
//  main.random()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:54 +0x15e
//  main.main.func1()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:34 +0x37
//
//Previous write at 0x00c4200b8048 by goroutine 6:
//  main.random()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:54 +0xcb
//  main.main.func1()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:34 +0x37
//
//Goroutine 7 (running) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:33 +0xb3
//
//Goroutine 6 (finished) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:33 +0xb3
//==================
//==================
//WARNING: DATA RACE
//Read at 0x00c4200b8048 by goroutine 8:
//  runtime.growslice()
//      /Users/alyson/.goenv/versions/1.10.3/src/runtime/slice.go:89 +0x0
//  main.random()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:54 +0x15e
//  main.main.func1()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:34 +0x37
//
//Previous write at 0x00c4200b8048 by goroutine 6:
//  main.random()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:54 +0xcb
//  main.main.func1()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:34 +0x37
//
//Goroutine 8 (running) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:33 +0xb3
//
//Goroutine 6 (finished) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/exercises/template1/template1.go:33 +0xb3
//==================
//Found 3 data race(s)
//exit status 66


// after apply mutex on rand()
//go run -race template1.go > /dev/null
//0 66
//1 41
//2 44
//3 80
//4 19
//5 60
//6 58
//7 59
//8 16
//9 54
//10 62
//11 56
//12 12
//13 80
//14 81
//15 69
//16 4
//17 63
//18 35
//19 44
//20 58
//21 72
//22 6
//23 3
//24 38
//25 92
//26 3
//27 65
//28 88
//29 95
