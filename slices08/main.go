package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato"}
	fmt.Printf("Type of fruitList %T \n", fruitList)

	fruitList = append(fruitList, "Lund Pakora", "Jhant Mix")
	fmt.Println("Fruitlist", fruitList)
	fmt.Println("Length if fruitlist", len(fruitList))

	fruitList = append(fruitList[1:3], fruitList...) // last index is not inclusive i.e 3 is not inclusive
	fmt.Println(fruitList)

	highScores := make([]int, 4) // create a slice of size 4

	highScores[0] = 12
	highScores[1] = 4
	highScores[3] = 38

	highScores = append(highScores, 78, 39, 12) // even after setting the size append can work as it resizes the whole slice

	//fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores) // sorting slice using ints

	//fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove value from slices based on index

	var courses = []string{"lunth", "nagar", "ka", "don", "hu"}

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...) // remove value at index 2
	fmt.Println(courses)

	// for i := 0; i < len(courses); i++ {
	// 	fmt.Printf("Course val %v \n", courses[i])
	// }

	for j := range courses {
		fmt.Printf("Courses val %v \n", courses[j])
	}
}
