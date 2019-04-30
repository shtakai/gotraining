// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// pointer or value
// think about the nature of the type. pick one. be consistent.
// - if it represents something unique: pointer
//   (factory method)
// - if it is a basic type: value
// - if it is implemented with a reference type: value
//     - avoid double references

// notify implements a method with a value receiver.
func (u user) notify() {
	// receiver variable: u user
	// other language -> this/self
	// Copy of u
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	// pointer receiver
	// Actual u
	//   => modifier
	u.email = email
}

func main() {

	// Values of type user can be used to call methods
	// declared with both value and pointer receivers.
	bill := user{"Bill", "bill@email.com"}
	//bill.changeEmail("bill@hotmail.com") // NOT *user
	// compiler adjust this ^ to ↓
	(&bill).changeEmail("bill@hotmail.com") // NOT *user
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with both value and pointer receiver.
	joan := &user{"Joan", "joan@email.com"}
	//       address
	joan.changeEmail("joan@hotmail.com")
	//joan.notify()
	// compiler adjust this ^ to ↓
	(*joan).notify()

	//Sending User Email To Bill<bill@hotmail.com>
	//Sending User Email To Joan<joan@hotmail.com>


	// Create a slice of user values with two users.
	users := []user{  // []*user はだめ
		{"ed", "ed@email.com"},
		{"erick", "erick@email.com"},
	}

	// Iterate over the slice of users switching
	// semantics. Not Good!
	// これは動かない。
	//for _, u := range users {
	//	// u: copy
	//	u.changeEmail("it@wontmatter.com")
	//}
	//
	//fmt.Println(users)
	//[{ed ed@email.com} {erick erick@email.com}]

	// ↓ か []*userだが、後者は臭うのでだめ。
	// ↓がいい
	for i := range users {
		// u: copy
		users[i].changeEmail("it@wontmatter.com")
	}
	fmt.Println(users)
	//[{ed it@wontmatter.com} {erick it@wontmatter.com}]

	// 別の話だが、getter/setterはだめ。


	// Exception example: Using pointer semantics
	// for a collectoin of strings.
	keys := make([]string, 10)
	for i := range keys {
		keys[i] = func() string { return "key-gen" }()
	}
}
