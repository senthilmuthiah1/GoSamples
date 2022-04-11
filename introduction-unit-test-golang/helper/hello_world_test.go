package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Farhan1")
	if result != "Hello Farhan" {
		// error message
		t.Fatal("Result must be 'Hello Farhan")
	}
	fmt.Println("Done")
}
