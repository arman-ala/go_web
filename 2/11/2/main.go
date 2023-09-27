package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	names := []string{"Arman", "Soheil", "Dean", "Sam"}
	tpl := template.Must(template.ParseGlob("*.gohtml"))

	index, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.ExecuteTemplate(index, "index.gohtml", names)
	if err != nil {
		fmt.Println(err)
	}
}
