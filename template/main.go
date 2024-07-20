package main

import (
	"html/template"
	"log"
	"os"
)

type Course struct {
	Name       string
	TotalHours int
}

func main() {
	course := Course{"GO", 40}
	tmp := template.New("CourseTemplate")

	tmp, _ = tmp.Parse("Course: {{.Name}} - Total Hours: {{.TotalHours}}")

	//podemos passar qualquer coisa do tipo io.Writer
	err := tmp.Execute(os.Stdout, course)

	if err != nil {
		log.Fatal(err)
	}
}
