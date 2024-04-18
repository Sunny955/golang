package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday"}

	for i := 0; i < len(days); i++ {
		fmt.Println("Days here", days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v \n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		// if rougueValue == 5 {
		// 	break
		// }

		if rougueValue == 7 {
			rougueValue++
			continue
		}

		fmt.Println("Rougue value is", rougueValue)
		rougueValue++
	}
}
