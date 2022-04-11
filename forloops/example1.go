/* Foreach loop. Golang does not have a foreach keyword.
But we can use a foreach loop by receiving 2 values in a for-range loop.
 We ignore the first, and the second is each element's value*/

package main

import "fmt"

func main() {
	// Two strings in a slice.
	animals := []string{"bird", "frog"}

	// Loop over each element directly (foreach loop).
	// ... Ignore the first pair of each returned pair (the index).
	for _, animal := range animals {
		fmt.Println(animal)
	}
}
