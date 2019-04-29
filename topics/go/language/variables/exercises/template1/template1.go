// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var a int
	var b bool
	var c string

	// Display the value of those variables.
	fmt.Println("zero value")
	fmt.Printf("a int %v\n", a)
	fmt.Printf("b bool %v\n", b)
	fmt.Printf("c string %v\n", c)

	// := is used for new var
	// 1) var name string
	//    name = "Jacob"
	// 2) name := "Jacob"

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 1020
	bb := false
	cc := "sashimiudon"

	// Display the value of those variables.
	fmt.Println("init")
	fmt.Printf("aa int %v\n", aa)
	fmt.Printf("bb bool %v\n", bb)
	fmt.Printf("cc string %v\n", cc)

	// Perform a type conversion.
	aaa := int64(aa)

	// Display the value of that variable.
	fmt.Println("conv")
	fmt.Printf("aaa int64 %v\n", aaa)
}

// excercises
