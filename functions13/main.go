package main

import "fmt"

func main() {
	sum := adder(2, 4)
	proSum, message := proAdder(19, 28, 4, 12)
	func(msg string) {
		fmt.Println("Message:", msg)
	}("Hello")

	mult := func(a int, b int) int {
		return a * b
	}(7, 8)

	fmt.Println("Sum of a and b->", sum)
	fmt.Println("Multiply of a and b->", mult)
	fmt.Println(message, proSum)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0

	// here we treat values as slice
	for _, val := range values {
		total += val
	}
	return total, "Hello you got this value"
}
