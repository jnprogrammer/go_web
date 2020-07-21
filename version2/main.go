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

	newfile, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer newfile.Close()

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
