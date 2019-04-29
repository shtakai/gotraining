// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var slice []int

	// Append numbers to the slice.
	for i :=0; i < 10; i++ {
		slice = append(slice, i*10)
	}

	fmt.Println("Display each value in the slice.")
	// Display each value in the slice.
	for _, v := range slice { // _ : blank identifier
		fmt.Println(v)
	}

	// get index
	//for index := range slice { // _ : blank identifier

	// Declare a slice of strings and populate the slice with names.
    strSlice := []string {"one", "two", "three", "four", "five"}

	// Display each index position and slice value.
	fmt.Println("Display each index/value in the slice value.")
	for i, v := range strSlice {
		fmt.Println("index: ", i, " value: ",v)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	oneTwo := strSlice[1:3]

	// Display each index position and slice values for the new slice.
	fmt.Println("Display each index/value in the new slice value.")
	for i, v := range oneTwo{
		fmt.Println("index: ", i, " value: ",v)
	}
}
