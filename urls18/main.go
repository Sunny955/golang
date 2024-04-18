package main

import (
	"fmt"
	"net/url"
)

const Url string = "https://www.w3schools.com:3000/typescript/typescript_getstarted.php?mulund=jsdh&kotha=meenabazaar"

func main() {
	// parsing url data
	result, err := url.Parse(Url)

	if err != nil {
		panic(err)
	}

	// fmt.Println("Scheme of the url", result.Scheme)
	// fmt.Println("Host of the url", result.Host)
	// fmt.Println("Path of the url", result.Path)
	// fmt.Println("Port of the url", result.Port())
	// fmt.Println("Query inside the url", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T \n", qparams)           // the type is of map, key-value
	fmt.Printf("The type of query params are: %T \n", qparams["mulund"]) // slice type

	fmt.Println(qparams["mulund"])

	for key, val := range qparams {
		fmt.Printf("Key is %v value is %v \n", key, val)
	}

	partsOfUrl := &url.URL{
		Scheme:  result.Scheme,
		Host:    result.Host,
		Path:    result.Path,
		RawPath: result.RawPath,
	}

	anotherURL := partsOfUrl.String()
	fmt.Println("URL formed is->", anotherURL)
}
