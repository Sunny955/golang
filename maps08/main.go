package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string) // defining map key value pair

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Printf("The type of languages is: %T \n", languages)
	fmt.Println(languages)
	fmt.Println(languages["RB"])

	delete(languages, "RB")
	fmt.Println(languages["RB"])

	// loops thorugh the maps

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
