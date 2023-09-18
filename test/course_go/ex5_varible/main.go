package main

import (
	"html/template"
	"log"
	"os"
)

var (
	tpl   *template.Template
	Sages []sage
	Data  items
)

type person struct {
	First_name string
	Last_name  string
	Age        int
}
type sage struct {
	Name  string
	Motto string
}
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func data1() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	Sages = []sage{buddha, gandhi, mlk, jesus, muhammad}
}

func data2() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	Data = items{
		Wisdom:    sages,
		Transport: cars,
	}
}

func main() {
	// data
	text := "hello there"
	name := []string{"Bob", "Artyom", "Ilya", "Alex"}
	send1 := map[string]string{
		"India":     "Gandhi",
		"America":   "NYC",
		"Russia":    "Moscow",
		"Kazakhsan": "Astana",
	}
	send2 := person{
		First_name: "Artyom",
		Last_name:  "Gruber",
		Age:        25,
	}
	data1()
	data2()
	// string
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl1.html", text); err != nil {
		log.Fatalln(err)
	}
	// slice
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl2.html", name); err != nil {
		log.Fatalln(err)
	}
	// map
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl3.html", send1); err != nil {
		log.Fatalln(err)
	}
	// struct
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl4.html", send2); err != nil {
		log.Fatalln(err)
	}
	// slice struct
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl5.html", Sages); err != nil {
		log.Fatalln(err)
	}
	// struct slice struct
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl6.html", Data); err != nil {
		log.Fatalln(err)
	}
}
