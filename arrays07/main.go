package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Pineapple"
	fruitList[3] = "Cakeapple"

	fmt.Println("List of fruits-", fruitList)
	fmt.Println("Length of fruit-", len(fruitList))

	var veglist = [3]string{"Mote", "Mote", "Papite"}
	fmt.Println("Vegetables -", veglist)
	fmt.Println("Length of Vegetables -", len(veglist))
}
