package main

import "fmt"

func main() {

	// defer will run the line of code in the end of the program

	defer fmt.Println("One")   // printed 5th
	defer fmt.Println("Two")   // printed 4th
	defer fmt.Println("Three") // printed 3rd

	fmt.Println("Hello") // printed 1st

	myDefer() // printed 2nd
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
