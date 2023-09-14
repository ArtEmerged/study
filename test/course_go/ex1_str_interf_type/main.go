package main

import "fmt"

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

type person struct {
	fname string
	lname string
}
type SecretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good Morning, James."`)
}

func (sa SecretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

func main() {
	p1 := person{"Jon", "Conor"}

	sa1 := SecretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(p1)
	saySomething(sa1)
}
