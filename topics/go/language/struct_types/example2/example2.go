// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and initialize anonymous
// struct types.
package main

import "fmt"

func main() {

	// Declare a variable of an anonymous type set
	// to its zero value.
	// var e1 struct { ... }  <= アノニマスタイプ
	// 利点: one-off/one-shot situation
	// JSON document, api
	//  zero value
	var e1 struct {
		// zero value
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymous type and init
	// using a struct literal.
	e2 := struct { // := literal value { }
		// zero value => anonymous type
		flag    bool // no need ,
		counter int16
		pi      float32
	}{
		// literal value
		// non zero value
		flag:    true, // need : and ,
		counter: 10,
		pi:      3.141592,
	}

	// Display the values.
	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
