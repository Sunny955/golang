package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API handling!!!!")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{"101", "Fuck your fear", 289, &Author{"Kothan Gupta", "www.udemy.com"}})

	courses = append(courses, Course{"102", "Dalli Gym me aaun", 380, &Author{"Chuttan Lodha", "www.udacity.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello welcome to my server!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses!")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a single course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find the matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id!")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a single course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!")
	}

	// if body is there but data is empty {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data entered!")
		return
	}

	// generate unique id, string
	// append course into courses

	// rand.Seed(time.Now().UnixNano())
	// course.CourseId = strconv.Itoa(rand.Intn(100))
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomValue := rng.Intn(100)
	course.CourseId = strconv.Itoa(randomValue)

	if course.CourseName == "" {
		json.NewEncoder(w).Encode("Please enter the course name")
		return
	}

	if course.CoursePrice == 0 {
		json.NewEncoder(w).Encode("course price can't be 0 or null")
		return
	}

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a single course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop through courses, get the id ,remove and add with params[id]

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			err := json.NewDecoder(r.Body).Decode(&updatedCourse)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if updatedCourse.Author == nil {
				updatedCourse.Author = course.Author
			}

			if updatedCourse.CoursePrice == 0 {
				updatedCourse.CoursePrice = course.CoursePrice
			}

			if updatedCourse.CourseName == "" {
				updatedCourse.CourseName = course.CourseName
			}

			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("No course is found, can't update!")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a single course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, find id and remove it
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course with given id is deleted successfully!")
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found with the given id!")
	return
}
