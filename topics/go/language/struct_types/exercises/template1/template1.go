// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main


// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name string
	age  int8
	good bool
}

func main() {

	// Declare variable of type user and init using a struct literal.
	var sushi user
	sushix :=user {
		name: "sushiX",
		age: 2,
		good: true,
	}

	sashimi := user {
		name: "sashimi",
		age: 5,
		good: false,
	}

	// Display the field values.
	fmt.Println("")
	fmt.Println("sushi")
	fmt.Println(sushi)
	fmt.Println(sushi.name)
	fmt.Println(sushi.age)
	fmt.Println(sushi.good)
	fmt.Println("")
	fmt.Println("sushix")
	fmt.Println(sushix)
	fmt.Println(sushix.name)
	fmt.Println(sushix.age)
	fmt.Println(sushix.good)
	fmt.Println("")
	fmt.Println("sashimi")
	fmt.Println(sashimi)
	fmt.Println(sashimi.name)
	fmt.Println(sashimi.age)
	fmt.Println(sashimi.good)

	// Declare a variable using an anonymous struct.
	food := struct {
		name string
		age int8
		good bool
	}{
		name: "soba",
		age: 35,
		good: true,
	}

	// Display the field values.
	fmt.Println("")
	fmt.Println("anonymous")
	fmt.Println("food")
	fmt.Println(food)
	fmt.Println(food.name)
	fmt.Println(food.age)
	fmt.Println(food.good)
}
