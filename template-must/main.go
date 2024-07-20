package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name       string
	TotalHours int
}

func main() {
	course := Course{"GO", 40}

	t := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Total Hours: {{.TotalHours}}"))

	err := t.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
