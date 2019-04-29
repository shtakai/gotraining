// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how only types that can have
// equality defined on them can be a map key.
package main

import "fmt"

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

// users defines a set of users.
type users []user // storage: 1word  key type must be comparative < > =
// key is not func/map/slice
// ok struct/array/pointer

func main() {

	// Declare and make a map that uses a slice as the key.
	// create map for me!
	u := make(map[users]int)

	// ./example3.go:22: invalid map key type users
	//# command-line-arguments
	//./example3.go:22:12: invalid map key type users

	// Iterate over the map.
	for key, value := range u {
		fmt.Println(key, value)
	}
}
