package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, r.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
