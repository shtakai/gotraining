// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a program that declares two anonymous functions. One that counts down from
// 100 to 0 and one that counts up from 0 to 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

// Add imports.
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
		fmt.Println("i decrement start")
		// Declare a loop that counts down from 100 to 0 and
		// display each value.
		for i := 100; i >=0; i-- {
			fmt.Println("i decrement: ", i)
		}
		// Tell main we are done.
		fmt.Println("i decrement done")
		wg.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		fmt.Println("j increment start")
		// Declare a loop that counts up from 0 to 100 and
		// display each value.
		for j := 0; j <100; j++ {
			fmt.Println("j increment: ", j)
		}

		// Tell main we are done.
		fmt.Println("j increment done")
        wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("wait finishing goroutines")
	wg.Wait()

	// Display "Terminating Program".
	fmt.Println("Terminating Program")
}

//Start Goroutines
//wait finishing goroutines
//j increment start
//j increment:  0
//j increment:  1
//j increment:  2
//j increment:  3
//j increment:  4
//j increment:  5
//j increment:  6
//j increment:  7
//j increment:  8
//j increment:  9
//j increment:  10
//j increment:  11
//j increment:  12
//j increment:  13
//j increment:  14
//j increment:  15
//j increment:  16
//j increment:  17
//j increment:  18
//j increment:  19
//j increment:  20
//j increment:  21
//j increment:  22
//j increment:  23
//j increment:  24
//j increment:  25
//j increment:  26
//j increment:  27
//j increment:  28
//j increment:  29
//j increment:  30
//j increment:  31
//j increment:  32
//j increment:  33
//j increment:  34
//j increment:  35
//j increment:  36
//j increment:  37
//j increment:  38
//j increment:  39
//j increment:  40
//j increment:  41
//j increment:  42
//j increment:  43
//j increment:  44
//j increment:  45
//j increment:  46
//j increment:  47
//j increment:  48
//j increment:  49
//j increment:  50
//j increment:  51
//j increment:  52
//j increment:  53
//j increment:  54
//j increment:  55
//j increment:  56
//j increment:  57
//j increment:  58
//j increment:  59
//j increment:  60
//j increment:  61
//j increment:  62
//j increment:  63
//j increment:  64
//j increment:  65
//j increment:  66
//j increment:  67
//j increment:  68
//j increment:  69
//j increment:  70
//j increment:  71
//j increment:  72
//j increment:  73
//j increment:  74
//j increment:  75
//j increment:  76
//j increment:  77
//j increment:  78
//j increment:  79
//j increment:  80
//j increment:  81
//j increment:  82
//j increment:  83
//j increment:  84
//j increment:  85
//j increment:  86
//j increment:  87
//j increment:  88
//j increment:  89
//j increment:  90
//j increment:  91
//j increment:  92
//j increment:  93
//j increment:  94
//j increment:  95
//j increment:  96
//j increment:  97
//j increment:  98
//j increment:  99
//j increment done
//i decrement start
//i decrement:  100
//i decrement:  99
//i decrement:  98
//i decrement:  97
//i decrement:  96
//i decrement:  95
//i decrement:  94
//i decrement:  93
//i decrement:  92
//i decrement:  91
//i decrement:  90
//i decrement:  89
//i decrement:  88
//i decrement:  87
//i decrement:  86
//i decrement:  85
//i decrement:  84
//i decrement:  83
//i decrement:  82
//i decrement:  81
//i decrement:  80
//i decrement:  79
//i decrement:  78
//i decrement:  77
//i decrement:  76
//i decrement:  75
//i decrement:  74
//i decrement:  73
//i decrement:  72
//i decrement:  71
//i decrement:  70
//i decrement:  69
//i decrement:  68
//i decrement:  67
//i decrement:  66
//i decrement:  65
//i decrement:  64
//i decrement:  63
//i decrement:  62
//i decrement:  61
//i decrement:  60
//i decrement:  59
//i decrement:  58
//i decrement:  57
//i decrement:  56
//i decrement:  55
//i decrement:  54
//i decrement:  53
//i decrement:  52
//i decrement:  51
//i decrement:  50
//i decrement:  49
//i decrement:  48
//i decrement:  47
//i decrement:  46
//i decrement:  45
//i decrement:  44
//i decrement:  43
//i decrement:  42
//i decrement:  41
//i decrement:  40
//i decrement:  39
//i decrement:  38
//i decrement:  37
//i decrement:  36
//i decrement:  35
//i decrement:  34
//i decrement:  33
//i decrement:  32
//i decrement:  31
//i decrement:  30
//i decrement:  29
//i decrement:  28
//i decrement:  27
//i decrement:  26
//i decrement:  25
//i decrement:  24
//i decrement:  23
//i decrement:  22
//i decrement:  21
//i decrement:  20
//i decrement:  19
//i decrement:  18
//i decrement:  17
//i decrement:  16
//i decrement:  15
//i decrement:  14
//i decrement:  13
//i decrement:  12
//i decrement:  11
//i decrement:  10
//i decrement:  9
//i decrement:  8
//i decrement:  7
//i decrement:  6
//i decrement:  5
//i decrement:  4
//i decrement:  3
//i decrement:  2
//i decrement:  1
//i decrement:  0
//i decrement done
//Terminating Program
