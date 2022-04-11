// Golang program to demonstrate the declaration
// and initialization of pointers
package main

func main() {

	var num1 int = 100
	println("The value is of num1 is ", num1)
	println("The Address  of num1 is ", &num1)

	// Let us store the address

	var ptrNum1 *int
	ptrNum1 = &num1

	println("The value is of ptrNum1 is ", ptrNum1)
	println("The value is of *ptrNum1 is ", *ptrNum1)

	num2 := 200
	println("The value is of num2 is ", num2)
	println("The Address  of num2 is ", &num2)

	// Let us store the address
	ptrNum2 := &num2

	println("The value is of ptrNum2 is ", ptrNum2)
	println("The value is of *ptrNum2 is ", *ptrNum2)

}
