// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and initialize struct types.
package main

import "fmt"

// example represents a type with different fields.
// type で宣言
// name: example
// struct => structure (composite type)
type example struct {
	flag    bool// false
	counter int16 // 0
	pi      float32 // 0
}

// |flag|counter|pi      |
//   1B   2B      4B
// flag    1B (start 0
// counter 2B (start 1 ??????!!! => starts w/padding+1B
// pi      4B
// -------------
// total   7B
//
//  => add padding
// |flag| p |counter|pi      |
//   1B  1B   2B      4B

//              開始可能
// 1B  -=-=-=-= 0 1 2 3 4 5 6 7
// 2B  --==--== 0   2   4   6
// 4B  ----==== 0       4
// 8B  -------- 0
// 開始と終了バイトの格納ルールに注意。外れるならpadding入れる

//   1B 1B 2B    4B
//   f  p  count pi
//   0  1  2  3  4567

//  bool +       float32
//  1B           4B
//  1B  P  P  P  4B------->
//  bl  p  p  p  pi------->
//  0   1  2  3  4 5 6 7

// example1: Mach-O 64-bit executable x86_64
func main() {

	// Declare a variable of type example set to its
	// zero value.
	var e1 example
	//  ^ name
	//         ^ struct type <= example (zero value

	// Display the value.
	// {flag:false counter:0 pi:0}
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using
	// a struct literal.
	// :=  literal value  {  }
	e2 := example{
		flag:    true, // : と , 入れる
		counter: 10,
		pi:      3.141592, // ラストにも : と , 入れる
	}

	// Display the field values.
	// access => use . (dot) like javascript
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
