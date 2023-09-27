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
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
}

/*
at the terminal:
go run main.go Arman
*/
