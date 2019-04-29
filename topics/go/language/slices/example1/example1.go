// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the capacity of the slice
// is not available for use.
// slice != array
package main

import "fmt"

func main() {
	// type []string
	// represent: ordered list of strings
	// stodage 3 words
	// pointer *
	//  length 5
	// capacity 5

	// *
	// 5
	// 5

	// *
	// 5
	// 5     x x x x    x
	//       |<length>  |
	//       |<capacity>|

	// Create a slice with a length of 5 elements.
	fruits := make([]string, 5)
	//  make(type[], length)   length: enough size
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// You can't access an index of a slice beyond its length.
	fruits[5] = "Runtime error"

	// Error: panic: runtime error: index out of range

	fmt.Println(fruits)
}
