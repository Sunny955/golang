package main

import (
	"fmt"
	"time"
)

// func main() {
// 	go greeter("Hello")
// 	greeter("World")
// }

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond) // Simulate some work
	}
}

func main() {
	go sayHello()                      // Start a new goroutine
	time.Sleep(500 * time.Millisecond) // Give some time for the goroutine to finish
	fmt.Println("Main function")
}

/*
In this example:

We define a function sayHello() which prints "Hello" five times with a small delay between each print.
In the main() function, we start a new goroutine by calling go sayHello().
Without the time.Sleep() call in main(),
the program might exit before the goroutine has finished executing,
so we add a delay to ensure the goroutine has time to complete.

*/
