package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sunny955/mongoapi/router"
)

func main() {
	fmt.Println("Database creating, updating and deleting APIs")
	r := router.Router()
	fmt.Println("Server is getting start.....")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is running at port 4000...")
}
