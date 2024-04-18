package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // don't show any thing to the user
	Tags     []string `json:"tags,omitempty"` // omit if this tag field is empty
}

func main() {
	// EncodeJson()
	// DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"React JS", 29, "Udemy", "1234", []string{"web-dev", "js"}},
		{"Fuckers", 1930, "Udemy", "qwerty", []string{"intercourse", "bg"}},
		{"Chuppa Laga", 386, "Pornhub", "1736", nil},
	}

	// converting the above to json
	// finalJson, err := json.Marshal(myCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println("final json data ->", string(finalJson))

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Fuckers",
		"price": 1930,
		"website": "Udemy",
		"tags": ["intercourse", "bg"]
	}
	`)

	var checkCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON data was valid")
		json.Unmarshal(jsonDataFromWeb, &checkCourse) // don't want to have copy of data so pass by refrence
		fmt.Printf("%#v \n", checkCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID!")
	}

	// in some cases you want to add data to key value pair
	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)
	// fmt.Printf("%#v \n", onlineData)

	for k, v := range onlineData {
		fmt.Printf("Key is %v , value is %v and type is: %T \n", k, v, v)
	}

}
