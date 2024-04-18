package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	diceNum := rand.Intn(6) + 1

	fmt.Println("random number", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spaces")
	case 3:
		fmt.Println("You can move 3 spaces")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spaces")
	case 5:
		fmt.Println("You can move 5 spaces")
		fallthrough
	case 6:
		fmt.Println("Dice value is 6 and you can open")
	default:
		fmt.Println("What the fuck was that?")
	}

}
