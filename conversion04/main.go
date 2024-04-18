package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for the pizza:")

	//comma ok || err ok
	input, _ := reader.ReadString('\n')

	// numRating, err := strconv.ParseFloat(input,64); gives conversion error as we are getting \n at the end of the string
	var numRating, err = strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Got error!", err)
	} else {
		fmt.Println("Added +1 to your rating", numRating+1)
	}
}
