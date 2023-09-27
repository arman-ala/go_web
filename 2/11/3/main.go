package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	people := map[string]int{"Arman": 21, "Soheil": 21, "Dean": 45, "Sam": 39}
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
