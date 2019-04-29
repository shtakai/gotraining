// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the basic concept of pass by value.
package main

func main() {

	// Declare variable of type int with a value of 10.
	count := 10 // t:int v:10

	// Display the "value of" and "address of" count.
	// &count => address of    & => address of
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	// 11 になるだろう....
	//increment(count) // これだと 10 ー＞ １１ー＞１０になる！？！？
	//          value
	//increment(&count) //コンパイルできない。 値で渡さないといけない!!!!
	//        address
	//./example1.go:18:12: cannot use &count (type *int) as type int in argument to increment

	increment(&count) // アドレスで渡して

	// ?????!!?!?!?!?!???! 11ではない
	// passing the body
	//count:  Value Of[ 10 ]  Addr Of[ 0xc42005bf70 ]
	//inc:    Value Of[ 11 ]  Addr Of[ 0xc42005bf60 ] 渡した先は違う
	//count:  Value Of[ 10 ]  Addr Of[ 0xc42005bf70 ]

	// * and & => OK!
	//count:  Value Of[ 10 ]  Addr Of[ 0xc42005bf70 ]
	//inc:    Value Of[ 0xc42005bf70 ]        Addr Of[ 0xc42005bf60 ]
	//count:  Value Of[ 11 ]  Addr Of[ 0xc42005bf70 ]

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
//go:noinline
//func increment(inc int) {
func increment(inc *int) { // ポインター(type)で受ける *=> type definition

	// Increment the "value of" inc.
	//inc++ // VALUE OF
	*inc++ // VALUE OF   * =>dereference

	//n := *inc // deference
	//n ++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
	//println("inc:\tValue Of[", n, "]\tAddr Of[", &n, "]")
}

// pointer
// rule

// &
// 1. & gives you the address of a var

// *
// 2. for every type T another type exists *T
//              float32                *float32
// 3. * dereferences a pointer.  gets the value a pointer points to
