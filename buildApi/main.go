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
type Author struct{
	Name string `json:"name"`;
	Website string `json:"website"`
}

type Course struct{
	CourseName string `json:"coursename"`;
	CourseId string `json:"courseid"`;
	CoursePrice int `json:"price"`;
	Author *Author `json:"author"`
}
var courses []Course;


func main() {
	fmt.Println("string with building api");

	r:=mux.NewRouter();

	courses=append(courses, Course{CourseName: "ReactJs",CourseId: "15",CoursePrice: 1520,Author: &Author{Name: "vipin",Website: "vipinnotes"}});
	courses=append(courses, Course{CourseName: "MERN",CourseId: "58",CoursePrice: 5000,Author: &Author{Name: "vipin",Website: "academify"}})

	r.HandleFunc("/",serveHome).Methods("GET");
	r.HandleFunc("/courses",getAllCourses).Methods("GET");
	r.HandleFunc("/course/{id}",getCourseById).Methods("GET");
	r.HandleFunc("/course",createOneCourse).Methods("POST");
	r.HandleFunc("/course/{id}",updateCourseById).Methods("PUT");
	r.HandleFunc("/course/{id}",deleteCourseById).Methods("DELETE");

	log.Fatal(http.ListenAndServe(":4006",r));
}




func (c *Course) IsEmpty() bool  {
	return c.CourseName=="" ;
}


func serveHome (w http.ResponseWriter , r *http.Request){

	w.Write([]byte("<h1>Header 1</h1><h1>Header 2</h1><h1>Header 3</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("sending all courses");
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request){
	fmt.Println("sending course by id");
	w.Header().Set("content-Type","application/json")

	params:=mux.Vars(r);

	for _,course:=range courses{
		if course.CourseId==params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found by this id : ")
}

func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("creating one course");
	w.Header().Set("content-Type","application/json")

	if r.Body==nil{
		json.NewEncoder(w).Encode("no data found")
		return ;
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course);

	if course.IsEmpty(){
		json.NewEncoder(w).Encode("please provide name")
		return ;
	}
	
	rand.Seed(time.Now().UnixNano());
	course.CourseId=strconv.Itoa(rand.Intn(100));
	courses=append(courses, course);
	json.NewEncoder(w).Encode(course)
}

func updateCourseById(w http.ResponseWriter,r *http.Request){
	fmt.Println("updating course by id");
	w.Header().Set("content-Type","application/json")
	params:=mux.Vars(r);

	for index,course:=range courses{
		if(course.CourseId==params["id"]){
			courses=append(courses[:index],courses[index+1:]...);
			var course Course;
			json.NewDecoder(r.Body).Decode(&course);
			course.CourseId=params["id"]
			courses=append(courses, course);
			json.NewEncoder(w).Encode(course)
			return;
		}
	}

	json.NewEncoder(w).Encode("no course found by this id");
}                               

func deleteCourseById(w http.ResponseWriter,r *http.Request){
	
	fmt.Println("deleting course by id");
	w.Header().Set("content-Type","application/json")

	params:=mux.Vars(r);

	for index,course:=range courses{
		if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted");
			return ;

		}
	}
	json.NewEncoder(w).Encode("No course found by this id")
}