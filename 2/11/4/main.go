package main

import (
	"fmt"
	"os"
	"text/template"
)

type person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	people := []person{
		{"Arman", 21, "male"},
		{"Soheil", 21, "male"},
		{"Dean", 45, "male"},
		{"Sam", 39, "male"},
	}
	tpl := template.Must(template.ParseGlob("*.gohtml"))

	index, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.ExecuteTemplate(index, "index.gohtml", people)
	if err != nil {
		fmt.Println(err)
	}
}
