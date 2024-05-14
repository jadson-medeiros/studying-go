package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name string
	Hours int
}

func main() {
	course := Course{"Go", 40}
	tmp := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Hours: {{.Hours}}"))
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}