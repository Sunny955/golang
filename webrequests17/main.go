package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://pkg.go.dev/golang.org/x/tools/cmd/godoc"

// const url string = "https://lco.dev"

func main() {
	fmt.Println("Web requests!")

	data, err := http.Get(url)

	if err != nil {
		panic(err)
	} else {
		defer data.Body.Close() // caller's responsibility to close the connection at the end
		fmt.Printf("Response is of type %T \n", data)

		databytes, err := io.ReadAll(data.Body) // don't we ioutil instead use io

		if err != nil {
			panic(err)
		}

		content := string(databytes)
		fmt.Println("Content of url->", content)
	}

}
