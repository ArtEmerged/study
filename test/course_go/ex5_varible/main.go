package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	text := "hello there"
	name:= []string {"Bob", "Artyom", "Ilya", "Alex"}
	if err := tpl.ExecuteTemplate(os.Stdout,"tpl1.html", text); err != nil {
		log.Fatalln(err)
	}
	//
	if err := tpl.ExecuteTemplate(os.Stdout,"tpl2.html", name); err != nil {
		log.Fatalln(err)
	}
}
