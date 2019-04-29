// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

// Add imports.

func main() {

	// Declare and make a map of integer type values.
	var values map[string]int
	//values := make(map[string]int)

	// Initialize some data into the map.
	values = map[string]int {
		"one": 1,
		"two": 2,
		"six": 6,
		"ten": 10,
	}

	// Display each key/value pair.
    for k, v := range values {
    	fmt.Println("key:", k, " val:", v)
	}
}
//key: one  val: 1
//key: two  val: 2
//key: ten  val: 10
//key: six  val: 6
