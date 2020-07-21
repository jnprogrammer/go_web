package main

import (
	"log"
	"os"
	"text/template"
)

//ptr to a template is a container is to which all the templates get parsed and are held.
func main() {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
