package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("all.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	data := "INDEX PUGE"
	tpl.Execute(w, data)
}

func dog(w http.ResponseWriter, r *http.Request) {
	data := "DOG PUGE"
	tpl.Execute(w, data)
}

func me(w http.ResponseWriter, r *http.Request) {
	data := "MY NAME IS ARTYOM PUGE"
	err := tpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
