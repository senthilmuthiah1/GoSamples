package main

import "fmt"

type Claims struct {
	ClaimNumber string
	BatchName   string
	bookmarks   []string
}

func main() {
	fmt.Println("Welcome to Struct ")
	myclaims := []Claims{
		{"PC101", "Batch101", []string{"Denied,OutOfNetwork"}},
		{"PC102", "Batch101", []string{"Denied,onHold"}},
	}

	fmt.Printf("%s", myclaims)
}
