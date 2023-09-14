package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	// tpl.Execute(os.Stdout,  nil)
	two, _ := tpl.ParseFiles("two.gohtml")
	two.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	two.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	tpl2, err := template.ParseGlob("template/")
	if err != nil {
		log.Fatal(err)
	}
	tpl2.Execute(os.Stdout, nil)

}
