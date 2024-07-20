package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name       string
	TotalHours int
}

type Courses []Course

func main() {

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Courses{
		{"GO", 40},
		{"JAVA", 40},
	})

	if err != nil {
		panic(err)
	}
}
