package main

import (
	. "log"
	. "os"
	t "text/template"
)

var tpl *t.Template

func init() {

	tpl = t.Must(t.ParseGlob("templates/"))
}

func main() {

	err := tpl.Execute(Stdout, nil)
	if err != nil {
		Fatalln(err)
	}

	err = tpl.ExecuteTemplate(Stdout, "vespa.gohtml", nil)
	if err != nil {
		Fatalln(err)
	}

	err = tpl.ExecuteTemplate(Stdout, "vespa.gohtml", nil)
	if err != nil {
		Fatalln(err)
	}

	err = tpl.ExecuteTemplate(Stdout, "one.gohtml", nil)
	if err != nil {
		Fatalln(err)
	}
}
