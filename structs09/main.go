package main

import "fmt"

func main() {
	// there is no inheritence in golang
	// Also no super or parent concept

	sunny := User{"Sunny", "sunny@gmail.com", "admin", 23}

	fmt.Println(sunny)
	fmt.Printf("Sunny details are: %+v\n", sunny)
	fmt.Printf("Name is %v, Email is %v \n", sunny.Name, sunny.Email)
}

type User struct {
	Name   string
	Email  string
	Status string
	Age    int
}
