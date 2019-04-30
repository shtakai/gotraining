// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and iterate over
// arrays of different types.

// list of values
package main

import "fmt"

func main() {
	// type [5]string
	// storage
	// nil nil nil nil nil
	//  0   0   0   0   0
	//--------------------
	//i 0   1   2   3   4

	//  *   *   *   *   *
	//  5   6   6   5   4
	//
	//  V   V   V   V   v
	//  Apple
	//      Orange
	//            ......

	// Declare an array of five strings that is initialized
	// to its zero value.
	var fruits [5]string // zero values
	//         num  type => array of 5 string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"
	//bounce check

	//fruits[5] = "Plum"
	//Invalid array index 5 (out of bounds for 5-element array)

	// Iterate over the array of strings.
	// range: iterate
	// i:index(int) 0 1 2 ,,,,
	// fruit: copy of each values "Apple",,,,
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// Declare an array of 4 integers that is initialized
	// with some values.
	//  := type (not zero value)
	//    [n]type{x, xx, xxx...}
	numbers := [4]int{10, 20, 30, 40}

	// Iterate over the array of numbers.
	// len (built in function) ===> about 10 built int functions
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
		//                    index
	}
}
