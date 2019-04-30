// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to walk through a map by
// alphabetical key order.
package main

import (
	"fmt"
	"sort"
)

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
		"Jackson": {"Michael", "Jackson"},
	}

	// Pull the keys from the map.
	var keys []string
	for key := range users {
		keys = append(keys, key)
	}

	fmt.Println("before sort(just grabbed)")
	for _, key := range keys {
		fmt.Println(key, users[key])
	}

	// Sort the keys alphabetically.
	sort.Strings(keys) //Package sort
	fmt.Println("after sort")

	// Walk through the keys and pull each value from the map.
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
	// 並びはcostがかかる iteration sort....
}

//before sort(just grabbed)
//Roy {Rob Roy}
//Ford {Henry Ford}
//Mouse {Mickey Mouse}
//Jackson {Michael Jackson}

//after sort
//Ford {Henry Ford}
//Jackson {Michael Jackson}
//Mouse {Mickey Mouse}
//Roy {Rob Roy}
