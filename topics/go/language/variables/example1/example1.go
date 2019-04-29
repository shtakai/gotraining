// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// The playground is actually a 64-bit env with 32-bit pointers
// The os/arch combo is named nacl/amd64p32

// Sample program to show how to declare variables.
package main

import "fmt"

func main() {

	// Declare variables that are set to their zero value.
	// all bit are set to 0 => zero value

	// r: representation s:storage

	var a int // int (64b)
	// r: 0
	// s: 32b 000...000 (based on platform)
	//    64b 000...000
	// portablitiy/ performance issues
	// int8 uint8 int64.....
	// int > convenient way
	// ex)
	// var a int        int [0]
	// var a int64        int64 [0]
	// int / int64 => readability problem => use int

	var b string // r: ""  s: 2words(8B/16B)
	// Pointer -> nil(actual bytes)
	// length  ->  0

	var c float64 // r:0 s 64b 000....000 64 bits
	var d bool // 1B 00000000

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10 // t:int r:10 s: 64b
	bb := "hello"
	//
	// w words 8B /16B
	// P ->  *  -> h e l l o
	// L ->  5(length)
	//
	cc := 3.14159 // should be float64
	dd := true // type is right hand. t:bool s:00000001 r: true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Specify type and perform a conversion.
	aaa := int32(10) // type: int32(type conv /not func)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}

/*
	Zero Values:

	Type Initialized Value
	Boolean false
	Integer 0
	Floating Point 0
	Complex 0i
	String "" (empty string)
	Pointer nil
*/
