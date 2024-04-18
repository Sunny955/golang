package main

import "fmt"

const LoginToken string = "djsdhsds" // using capital L here which makes it public
// and can be accessible to whole program

func main() {
	var username string = "hello"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoogedin bool = false
	fmt.Println(isLoogedin)
	fmt.Printf("Variable is of type: %T \n", isLoogedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
