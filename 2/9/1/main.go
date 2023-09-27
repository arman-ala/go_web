package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseGlob("*.gohtml"))

	index, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.ExecuteTemplate(index, "index.gohtml", "Soheil big heart!")
	if err != nil {
		fmt.Println(err)
	}
}
