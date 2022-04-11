// Golang program to demonstrate
// the use of type inference in
// Pointer variables

/* If you are specifying the data type along with the pointer declaration then the pointer will
   be able to handle the memory address of that specified data type variable. For example,
   if you taking a pointer of string type then the address of the variable that you
   will give to a pointer will be only of string data type variable, not any other type. */

/* To overcome the above mention problem you can use the Type Inference concept of the
   var keyword. There is no need to specify the data type during the declaration.
   The type of a pointer variable can also be determined by the compiler like a normal
   variable. Here we will not use the * operator. It will internally determine by the
   compiler as we are initializing the variable with the address of another variable.
*/

package main

import "fmt"

func main() {

	// using var keyword
	// we are not defining
	// any type with variable
	var y = 458

	// taking a pointer variable using
	// var keyword without specifying
	// the type
	var p = &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)
}
