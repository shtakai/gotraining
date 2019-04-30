// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// $ ./example2 | cut -c1 | grep '[AB]' | uniq

// Sample program to show how the goroutine scheduler
// will time slice goroutines on a single thread.
package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"

	//"runtime"
	"strconv"
	"sync"
)

// init():  run before main()
func init() {

	// Allocate one logical processor for the scheduler to use.
	// => affect whole package
	// allow upto 1 process
	runtime.GOMAXPROCS(1)

	// if erase: runtime.GOMAXPROCS(1)
	// almost parallel
	//B: 49963: 02379422150eca886fdc1cd35627861091ef086f
	//A: 49927: 7b3733b4419ce766c0e7c3c2bcfdd2d90b436e73
	//A: 49928: ec22a5b4584ff9eebdd25e4496595e33bd14b6f4
	//A: 49929: c30fa68716d5f99f37e3f9b920bb61c3456678ed
	//A: 49930: eb753472fb36d5ea6798d870a059442be53037d5
	//B: 49964: 028a1c3b8b4f02d01d98f7e174d083928f3fd944
	//B: 49965: 2aa0d3a8dacae3359f68d61db6dad3da6639c9f4
}

func main() {
	//runtime.GOMAXPROCS(1)

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)
	// goroutinesがかわらないとき
	// goroutineが入る前ににwg.Add(1)する

	fmt.Println("Create Goroutines")

	// Create the first goroutine and manage its lifecycle here.
	//wg.Add(1) =>ここにいれるのはOK oooooooooo
	go func() {
		//wg.Add(1) => ここにいれるのはだめ   何もしない XXXXXXXXXX
		printHashes("A")
		wg.Done()
	}()

	// Create the second goroutine and manage its lifecycle here.
	//wg.Add(1)
	go func() {
		printHashes("B")
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printHashes calculates the sha1 hash for a range of
// numbers and prints each in hex encoding.
func printHashes(prefix string) {

	// print each has from 1 to 10. Change this to 50000 and
	// see how the scheduler behaves.
	for i := 1; i <= 50000; i++ {

		// Convert i to a string.
		num := strconv.Itoa(i)

		// Calculate hash for string num.
		sum := sha1.Sum([]byte(num))

		// Print prefix: 5-digit-number: hex encoded hash
		fmt.Printf("%s: %05d: %x\n", prefix, i, sum)
	}

	fmt.Println("Completed", prefix)
}
// 10
//Create Goroutines
//Waiting To Finish
//B: 00001: 356a192b7913b04c54574d18c28d46e6395428ab
//B: 00002: da4b9237bacccdf19c0760cab7aec4a8359010b0
//B: 00003: 77de68daecd823babbb58edb1c8e14d7106e83bb
//B: 00004: 1b6453892473a467d07372d45eb05abc2031647a
//B: 00005: ac3478d69a3c81fa62e60f5c3696165a4e5e6ac4
//B: 00006: c1dfd96eea8cc2b62785275bca38ac261256e278
//B: 00007: 902ba3cda1883801594b6e1b452790cc53948fda
//B: 00008: fe5dbbcea5ce7e2988b8c69bcfdfde8904aabc1f
//B: 00009: 0ade7c2cf97f75d009975f4d720d1fa6c19f4897
//B: 00010: b1d5781111d84f7b3fe45a0852e59758cd7a87e5
//Completed B
//A: 00001: 356a192b7913b04c54574d18c28d46e6395428ab
//A: 00002: da4b9237bacccdf19c0760cab7aec4a8359010b0
//A: 00003: 77de68daecd823babbb58edb1c8e14d7106e83bb
//A: 00004: 1b6453892473a467d07372d45eb05abc2031647a
//A: 00005: ac3478d69a3c81fa62e60f5c3696165a4e5e6ac4
//A: 00006: c1dfd96eea8cc2b62785275bca38ac261256e278
//A: 00007: 902ba3cda1883801594b6e1b452790cc53948fda
//A: 00008: fe5dbbcea5ce7e2988b8c69bcfdfde8904aabc1f
//A: 00009: 0ade7c2cf97f75d009975f4d720d1fa6c19f4897
//A: 00010: b1d5781111d84f7b3fe45a0852e59758cd7a87e5
//Completed A
//Terminating Program

//50000
//» go run example2.go|grep 'Completed' -A 3 -B 3
//A: 49998: ba0d018a03381a964f242be4ebabd82a04d642e7
//A: 49999: b01f85c46b2ff221fcb313f8e5215daa82600a3b
//A: 50000: c2d4c5452f59cff5973dd9d08df95f8b54cad995
//Completed A
//B: 20208: b4863ebf077c944e9b681bc3df757b8c6c925da6
//B: 20209: b6599c24593290529806bfb0cc4f7429412100ec
//B: 20210: 666687f661e8e78db6ba55dfb2ad5e67ef1d5f4b
//--
//--
//B: 49998: ba0d018a03381a964f242be4ebabd82a04d642e7
//B: 49999: b01f85c46b2ff221fcb313f8e5215daa82600a3b
//B: 50000: c2d4c5452f59cff5973dd9d08df95f8b54cad995
//Completed B
//Terminating Program

// GopherCon 2018: Kavya Joshi - The Scheduler Saga
//https://www.youtube.com/watch?v=YHRO5WQGh0k

// goroutine has cost.
//  at first few kb at stack
//  if 2M goroutines... => many costs


// call goroutine -> call goroutine -> call.......
//  it can
//   呼ばれたgoroutineが落ちると(panic).... ＞チャネルで処理せい。
//   それ以外は続く。
//
// limit of goroutines => depends hw/memory
