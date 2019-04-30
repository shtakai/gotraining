// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a program that declares two anonymous functions. One that counts down from
// 100 to 0 and one that counts up from 0 to 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {

	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Count down from 100 to 0.
		for count := 100; count >= 0; count-- {
			fmt.Printf("[A:%d]\n", count)
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Count up from 0 to 100.
		for count := 0; count <= 100; count++ {
			fmt.Printf("[B:%d]\n", count)
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	// Display "Terminating Program".
	fmt.Println("\nTerminating Program")
}
//Start Goroutines
//Waiting To Finish
//[B:0]
//[B:1]
//[B:2]
//[B:3]
//[B:4]
//[B:5]
//[B:6]
//[B:7]
//[B:8]
//[B:9]
//[B:10]
//[B:11]
//[B:12]
//[B:13]
//[B:14]
//[B:15]
//[B:16]
//[B:17]
//[B:18]
//[B:19]
//[B:20]
//[B:21]
//[B:22]
//[B:23]
//[B:24]
//[B:25]
//[B:26]
//[B:27]
//[B:28]
//[B:29]
//[B:30]
//[B:31]
//[B:32]
//[B:33]
//[B:34]
//[B:35]
//[B:36]
//[B:37]
//[B:38]
//[B:39]
//[B:40]
//[B:41]
//[B:42]
//[B:43]
//[B:44]
//[B:45]
//[B:46]
//[B:47]
//[B:48]
//[B:49]
//[B:50]
//[B:51]
//[B:52]
//[B:53]
//[B:54]
//[B:55]
//[B:56]
//[B:57]
//[B:58]
//[B:59]
//[B:60]
//[B:61]
//[B:62]
//[B:63]
//[B:64]
//[B:65]
//[B:66]
//[B:67]
//[B:68]
//[B:69]
//[B:70]
//[B:71]
//[B:72]
//[B:73]
//[B:74]
//[B:75]
//[B:76]
//[B:77]
//[B:78]
//[B:79]
//[B:80]
//[B:81]
//[B:82]
//[B:83]
//[B:84]
//[B:85]
//[B:86]
//[B:87]
//[B:88]
//[B:89]
//[B:90]
//[B:91]
//[B:92]
//[B:93]
//[B:94]
//[B:95]
//[B:96]
//[B:97]
//[B:98]
//[B:99]
//[B:100]
//[A:100]
//[A:99]
//[A:98]
//[A:97]
//[A:96]
//[A:95]
//[A:94]
//[A:93]
//[A:92]
//[A:91]
//[A:90]
//[A:89]
//[A:88]
//[A:87]
//[A:86]
//[A:85]
//[A:84]
//[A:83]
//[A:82]
//[A:81]
//[A:80]
//[A:79]
//[A:78]
//[A:77]
//[A:76]
//[A:75]
//[A:74]
//[A:73]
//[A:72]
//[A:71]
//[A:70]
//[A:69]
//[A:68]
//[A:67]
//[A:66]
//[A:65]
//[A:64]
//[A:63]
//[A:62]
//[A:61]
//[A:60]
//[A:59]
//[A:58]
//[A:57]
//[A:56]
//[A:55]
//[A:54]
//[A:53]
//[A:52]
//[A:51]
//[A:50]
//[A:49]
//[A:48]
//[A:47]
//[A:46]
//[A:45]
//[A:44]
//[A:43]
//[A:42]
//[A:41]
//[A:40]
//[A:39]
//[A:38]
//[A:37]
//[A:36]
//[A:35]
//[A:34]
//[A:33]
//[A:32]
//[A:31]
//[A:30]
//[A:29]
//[A:28]
//[A:27]
//[A:26]
//[A:25]
//[A:24]
//[A:23]
//[A:22]
//[A:21]
//[A:20]
//[A:19]
//[A:18]
//[A:17]
//[A:16]
//[A:15]
//[A:14]
//[A:13]
//[A:12]
//[A:11]
//[A:10]
//[A:9]
//[A:8]
//[A:7]
//[A:6]
//[A:5]
//[A:4]
//[A:3]
//[A:2]
//[A:1]
//[A:0]
//
//Terminating Program