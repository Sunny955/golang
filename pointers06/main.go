package main

import "fmt"

func main() {
	myNumber := 25

	var ptr = &myNumber
	// var pointer *int;

	fmt.Println("Pointer value", ptr)
	fmt.Println("Pointer number", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Pointer value", *ptr)
}
