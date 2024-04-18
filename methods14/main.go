package main

import "fmt"

func main() {
	// there is no inheritence in golang
	// Also no super or parent concept

	sunny := User{"Sunny", "sunny@gmail.com", "admin", 23, false}

	fmt.Println(sunny)
	fmt.Printf("Sunny details are: %+v\n", sunny)
	fmt.Printf("Name is %v, Email is %v \n", sunny.Name, sunny.Email)

	sunny.getStatus()
	sunny.changeMail()

	fmt.Println("New email of user", sunny.Email)
}

type User struct {
	Name     string
	Email    string
	Status   string
	Age      int
	loggedIn bool
}

func (u User) getStatus() {
	fmt.Println("User status :", u.loggedIn)
}

func (u User) changeMail() {
	u.Email = "test@dev.com" // here u is passed as the copy and not the actual address of it
	fmt.Println("New email of user is:", u.Email)
}
