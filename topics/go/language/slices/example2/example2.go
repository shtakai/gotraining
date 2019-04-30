// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the components of a slice. It has a
// length, capacity and the underlying array.
package main

import "fmt"

func main() {
	// type []string
	// represent: ordered list of strings
	// stodage 3 words
	// pointer *
	//  length 5
	// capacity 8

	// *
	// 5
	// 8     x x x x x  _ _ _
	//       |<length>|
	//       |<capacity----->|

	// Create a slice with a length of 5 elements and a capacity of 8.
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	inspectSlice(fruits)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] s lives at %p /   s[i] lives at %p %s\n",
			i,
			&s,
			&slice[i],
			s)
	}
}
