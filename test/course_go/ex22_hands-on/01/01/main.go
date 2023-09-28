package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "It's Index")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "It's Dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Your name: Artyom")
}
