package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

type Pet struct {
	Name   string
	Gender string
}

type Data struct {
	People []Person
	Pets   []Pet
}

func main() {
	people := []Person{
		{"Arman", 21, "male"},
		{"Soheil", 21, "male"},
		{"Dean", 45, "male"},
		{"Sam", 39, "male"},
	}

	pets := []Pet{
		{"Eskandari", "male"},
		{"PedarSag", "female"},
		{"Putin", "male"},
		{"Trump", "male"},
	}

	data := Data{
		people,
		pets,
	}

	tpl := template.Must(template.ParseGlob("*.gohtml"))

	index, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.ExecuteTemplate(index, "index.gohtml", data)
	if err != nil {
		fmt.Println(err)
	}
}
