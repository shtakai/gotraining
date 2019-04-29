// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare, initialize and iterate
// over a map. Shows how iterating over a map is random.
package main

import "fmt"

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

func main() {

	// Declare and initialize the map with values.
	users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"}, // ',' mandatory
	}
	//"Roy":     user{"Rob", "Roy"},
	// key        val

	// Iterate over the map printing each key and value.
	// UNORDERED
	// $watch -n1 go run  example4.go
	for key, value := range users {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Iterate over the map printing just the keys.
	// Notice the results are different.
	for key := range users {
		fmt.Println(key, users[key])
	}
}
//Mouse {Mickey Mouse}
//Jackson {Michael Jackson}
//Roy {Rob Roy}
//Ford {Henry Ford}
//
//Roy {Rob Roy}
//Ford {Henry Ford}
//Mouse {Mickey Mouse}
//Jackson {Michael Jackson}