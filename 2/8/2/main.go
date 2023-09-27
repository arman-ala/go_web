package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// fmt.Println("start")
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		fmt.Println(err)
	}
}

/*
at the terminal:
go run main.go Arman
*/
