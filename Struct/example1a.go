//  Defining a Struct in Go
package main

import "fmt"

type Fruit struct {
	name string
}

func main() {
	fmt.Println("Welcome to go Struct")

	//  zero-value Struct
	/* When a struct is declared but has not assigned any value, it is then a zero-valued
	struct. That means, all the fields will take their respective zero-values from
	their types. */
	var apple Fruit
	fmt.Println(apple)
}
