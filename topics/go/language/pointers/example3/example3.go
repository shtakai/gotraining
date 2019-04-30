// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the basic concept of using a pointer
// to share data.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	// zero value
	name   string
	email  string
	logins int
}

func main() {

	// Declare and initialize a variable named bill of type user.
	bill := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	//** We don't need to include all the fields when specifying field
	// names with a struct literal.

	// Pass the "address of" the bill value.
	display(&bill)//address

	// Pass the "address of" the logins field from within the bill value.
	increment(&bill.logins)

	// Pass the "address of" the bill value.
	display(&bill)
}

// increment declares logins as a pointer variable whose value is
// always an address and points to values of type int.
func increment(logins *int) {
	*logins++
	fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n\n", &logins, logins, *logins)
}

// display declares u as user pointer variable whose value is always an address
// and points to values of type user.
func display(u *user) { // *user => pointer of user
	fmt.Printf("%p\t%+v\n", u, *u) // *u dereference
	fmt.Printf("Name: %q Email: %q Logins: %d\n\n", u.name, u.email, u.logins)
	// v also
	fmt.Printf("Name: %q Email: %q Logins: %d\n\n", (*u).name, (*u).email, (*u).logins)
}
//
//0xc420090180    {name:Bill email:bill@ardanlabs.com logins:0}
//Name: "Bill" Email: "bill@ardanlabs.com" Logins: 0
//
//&logins[0xc4200a0020] logins[0xc4200901a0] *logins[1]
//
//0xc420090180    {name:Bill email:bill@ardanlabs.com logins:1}
//Name: "Bill" Email: "bill@ardanlabs.com" Logins: 1
//

