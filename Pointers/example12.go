/*You can also use the shorthand (:=) syntax to declare and initialize the pointer variables.
  The compiler will internally determine the variable is a pointer variable if we are passing
  the address of the variable using &(address) operator to it. */
// Golang program to demonstrate
// the use of shorthand syntax in
// Pointer variables
package main

import "fmt"

func main() {

	// using := operator to declare
	// and initialize the variable
	y := 458

	// taking a pointer variable using
	// := by assigning it with the
	// address of variable y
	p := &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)
}
