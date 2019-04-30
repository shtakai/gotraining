// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create goroutines and
// how the scheduler behaves.
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

	// wg is used to manage concurrency.
	var wg sync.WaitGroup // zero value   wg: keep track goroutine
	wg.Add(2) // create 2 goroutines

	fmt.Println("Start Goroutines")

	// Create a goroutine from the lowercase function.
	// anonymous function
	// call this function w/prefix:go
	//  program moves to under line(next line)
	go func() {
		lowercase()
		wg.Done() // wait group done  decrement counter
	}()

	// ↓ ↑ は順序わからない

	// Create a goroutine from the uppercase function.
	// also moves to down.
	go func() {
		uppercase()
		wg.Done() // wait group done  decrement counter
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	// important
	// main function is blocked til wg count is 0
	wg.Wait()
	// then move on!

	fmt.Println("\nTerminating Program")
}

// lowercase displays the set of lowercase letters three times.
func lowercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

// uppercase displays the set of uppercase letters three times.
func uppercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}
//Start Goroutines
//Waiting To Finish
//A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z
//Terminating Program


// concurrency traps
// 1. incomplete work. (main is done but other func is not done)
//  ex) go uppercace()
// "never start a goroutine without knowing how it will stop"

// ex) wg.Add(3) or exit without wg.Done() in goroutine.
//  => DEADLOCK
//   wg.Done()  --1
//   wg.Wait()  waiting for 0

