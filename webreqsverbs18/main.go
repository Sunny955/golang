package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest("https://reqres.in/api/userss?page=2")
	// PerformPostJsonRequest("https://reqres.in/api/users")
	PerformPostFormRequest("https://reqres.in/api/users") // using form data to send
}

func PerformGetRequest(url string) {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		panic("The response isn't resolved")
	}

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount here ->", byteCount)

	fmt.Println(responseString.String())
}

func PerformPostJsonRequest(url string) {
	// json payload
	requestBody := strings.NewReader(`
		{
			"name" : "Dalli",
			"job" : "Gym jaaun"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)

	responseString.Write(content)

	fmt.Println(responseString.String())
}

func PerformPostFormRequest(Url string) {
	data := url.Values{}

	data.Add("name", "Lunth")
	data.Add("job", "Borthel")

	response, err := http.PostForm(Url, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)

	responseString.Write(content)

	fmt.Println(responseString.String())

}
