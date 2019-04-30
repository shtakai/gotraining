// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how maps are not safe for concurrent use by default.
// The runtime will detect concurrent writes and panic.
package main

import (
	"fmt"
	"sync"
)

// scores holds values incremented by multiple goroutines.
var scores = make(map[string]int) // not safe

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			scores["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			scores["B"]++
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final scores:", scores)
}
// go run -race example5.go
//==================
//WARNING: DATA RACE
//Read at 0x00c4200a2180 by goroutine 7:
//  runtime.mapaccess1_faststr()
//      /Users/alyson/.goenv/versions/1.10.3/src/runtime/hashmap_fast.go:172 +0x0
//  main.main.func2()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example5/example5.go:29 +0x7b
//
//Previous write at 0x00c4200a2180 by goroutine 6:
//  runtime.mapassign_faststr()
//      /Users/alyson/.goenv/versions/1.10.3/src/runtime/hashmap_fast.go:694 +0x0
//  main.main.func1()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example5/example5.go:22 +0xdc
//
//Goroutine 7 (running) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example5/example5.go:27 +0xbb
//
//Goroutine 6 (running) created at:
//  main.main()
//      /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example5/example5.go:20 +0x99
//==================
//Final scores: map[A:1000 B:1000]
//Found 1 data race(s)
//exit status 66


//go run  example5.go                                            1 ↵
//fatal error: concurrent map writes
//
//goroutine 5 [running]:
//runtime.throw(0x10c16a1, 0x15)
//        /Users/alyson/.goenv/versions/1.10.3/src/runtime/panic.go:616 +0x81 fp=0xc420047f20 sp=0xc420047f00 pc=0x1026481
//runtime.mapassign_faststr(0x10a5820, 0xc42008a180, 0x10bf09f, 0x1, 0x1157340)
//        /Users/alyson/.goenv/versions/1.10.3/src/runtime/hashmap_fast.go:779 +0x3ce fp=0xc420047f90 sp=0xc420047f20 pc=0x100a1de
//main.main.func1(0xc4200180f0)
//        /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example5/example5.go:22 +0x9c fp=0xc420047fd8 sp=0xc420047f90 pc=0x108de7c
//runtime.goexit()
//        /Users/alyson/.goenv/versions/1.10.3/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc420047fe0 sp=0xc420047fd8 pc=0x104ddb1
//created by main.main
//        /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example5/example5.go:20 +0x6f
//
//goroutine 1 [semacquire]:
//sync.runtime_Semacquire(0xc4200180fc)
//        /Users/alyson/.goenv/versions/1.10.3/src/runtime/sema.go:56 +0x39
//sync.(*WaitGroup).Wait(0xc4200180f0)
//        /Users/alyson/.goenv/versions/1.10.3/src/sync/waitgroup.go:129 +0x72
//main.main()
//        /Users/alyson/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example5/example5.go:34 +0x9f
//exit status 2