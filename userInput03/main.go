package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for the pizza:")

	//comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating,", input)
	fmt.Printf("Type of this rating is %T \n", input)
}
