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

// Model for our course - File

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var courses []Course

// middleware , helper- FILE

func (c *Course) IsEmpty() bool {
	if c.CourseName == "" {
		return true
	}
	return false
}

func main() {
	fmt.Println("MY API ")
	r := mux.NewRouter()

	//seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Sahil Makeshwar", Website: "github.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN stack", CoursePrice: 499, Author: &Author{Fullname: "Sahil Makeshwar", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCources).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - FILE

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello welcome to my api</h1>"))
}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req

	params := mux.Vars(r)

	//loop through the couses , find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
		return
	}

	//what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("please send some data")
		return
	}

	// generate unique id, string
	// append the course into course

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id form req

	params := mux.Vars(r)

	//loop ,id , remove, add with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO send a response when id is not found;

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delte one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop , id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
