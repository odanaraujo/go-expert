package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name       string
	TotalHours int
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		t.Execute(w, Courses{
			{"GO", 40},
			{"JAVA", 40},
		})
	})
	http.ListenAndServe(":8282", nil)
}
