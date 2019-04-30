// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import (
	"fmt"
)

// Add imports.

func main() {

	// Declare an array of 5 strings set to its zero value.
	var fiveStrings [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	preFiveStrings := [5]string {"あああao", "いいいt", "tっっっｄh","fo","fi"}

	// Assign the populated array to the array of zero values.
	fiveStrings = [5]string{"1", "2", "3", "4", "5"}

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	fmt.Println("from zero value")
	for i, _ := range fiveStrings {
        fmt.Println("value:", fiveStrings[i], " address:", &fiveStrings[i])
	}
	fmt.Println("from populated")
	for i, _ := range preFiveStrings {
		fmt.Println("value:", preFiveStrings[i], " address:", &preFiveStrings[i])
	}
}
