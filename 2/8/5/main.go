package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseGlob("./*.gmao"))

	err := tpl.ExecuteTemplate(os.Stdout, "sample.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
