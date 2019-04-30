// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how arrays of different sizes are
// not of the same type.
package main

import "fmt"

func main() {

	// Declare an array of 5 integers that is initialized
	// to its zero value.
	var five [5]int // length: 5 [5]int
	// Declare an array of 4 integers that is initialized
	// with some values.
	four := [4]int{10, 20, 30, 40} // length: 4 [4]int

	// Assign one array to the other
	five = four
	// NOT:   [5]int = [4]int !!!!!
	// array size is also type
	// ./example2.go:21: cannot use four (type [4]int) as type [5]int in assignment

	// how to fix it....?

	fmt.Println(four)
	fmt.Println(five)
}

// Mechanical Sympathy
// working with the hardware instead if in spete of the hardware
//

// Henry petroski(2015)
// https://news.ycombinator.com/item?id=7667409
//The most amazing achievement of the computer software industry is its continuing cancellation of the steady and staggering gains made by the computer hardware industry.
//— Henry Petroski


// slice: view of a segment of a backing array

// https://golang.org/pkg/crypto/sha1/