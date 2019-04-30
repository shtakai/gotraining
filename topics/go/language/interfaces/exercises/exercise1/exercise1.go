// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
package main

import "fmt"

// speaker implements the voice of anyone.
type speaker interface {
	speak()
}

// english represents an english speaking person.
type english struct{}

// speak implements the speaker interface using a
// value receiver.
func (english) speak() { // value receiver
	fmt.Println("Hello World")
}

// chinese represents a chinese speaking person.
type chinese struct{}

// speak implements the speaker interface using a
// pointer receiver.
func (*chinese) speak() { // pointer receiver
	fmt.Println("你好世界")
}

func main() {

	// Declare a variable of the interface speaker type
	// set to its zero value.
	var sp speaker // zero value

	// Declare a variable of type english.
	var e english

	// Assign the english value to the speaker variable.
	sp = e

	// Call the speak method against the speaker variable.
	sp.speak() // (english).speak
	//Hello World

	// Declare a variable of type chinese.
	var c chinese

	// Assign the chinese pointer to the speaker variable.
	sp = &c // *chineseType

	// Call the speak method against the speaker variable.
	sp.speak()  // (&chinese).speak
	//你好世界

	// Call the sayHello function with new values and pointers
	// of english and chinese.
	sayHello(english{}) // english OK
	//Hello World
	sayHello(&english{}) // &english OK
	//Hello World
	sayHello(&chinese{}) // &chinese OK
	//你好世界

	// Why does this not work?
	//sayHello(chinese{}) // compiler says 'this is not allowed'
	//# command-line-arguments
	//./exercise1.go:77:18: cannot use chinese literal (type chinese) as type speaker in argument to sayHello:
	//        chinese does not implement speaker (speak method has pointer receiver)
}

// sayHello abstracts speaking functionality.
func sayHello(sp speaker) {
	sp.speak()
}
