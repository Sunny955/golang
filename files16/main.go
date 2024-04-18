package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This is we are writing in a file!"

	file, err := os.Create("./myLcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./myLcogofile.txt")
}

func readFile(fileName string) {
	// databyte, err := ioutil.ReadFile(fileName) deprecated

	dataByte, err := os.ReadFile(fileName)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("The content of the file: ", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
