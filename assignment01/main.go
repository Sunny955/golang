package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	data, err := fetchData("https://reqres.in/api/userss?page=2")

	if err != nil {
		panic(err)
	}

	// changing the above byte data to json data

	var responseData map[string]interface{}

	if err := json.Unmarshal(data, &responseData); err != nil {
		panic(err)
	}

	for k, v := range responseData {
		fmt.Printf("key is %v, value is %v \n", k, v)
	}
}

func fetchData(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}
