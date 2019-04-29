// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

func main() {

	// Declare and make a map of integer type values.
	//departments := make(map[string]int)
	var departments map[string]int //=> zero value / nil pointer

	//fmt.Println(len(departments), departments["x"])}
	//# command-line-arguments
	//./exercise1.go:17:2: syntax error: non-declaration statement outside function body

	if departments == nil {
		fmt.Println("create because of nil")
		departments = make(map[string]int)
	}
	fmt.Println(len(departments), departments["x"])

	// Initialize some data into the map.
	departments["IT"] = 20
	departments["Marketing"] = 15
	departments["Executives"] = 5
	departments["Sales"] = 50
	departments["Security"] = 8

	// Display each key/value pair.
	for key, value := range departments {
		fmt.Printf("Dept: %s People: %d\n", key, value)
	}
}

//create because of nil
//0 0
//Dept: IT People: 20
//Dept: Marketing People: 15
//Dept: Executives People: 5
//Dept: Sales People: 50
//Dept: Security People: 8e: 15