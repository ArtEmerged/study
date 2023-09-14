package main

import (
	"fmt"
	"math"
)

// type animal interface {
// 	Sound()
// 	Len()
// }

// type cat struct {
// 	count int
// }

// type dog struct {
// 	count int
// }

// func (c *cat) Sound() {
// 	n := c.count
// 	for n > 0 {
// 		fmt.Println("Мяу!")
// 		n--
// 	}
// }

// func (c *cat) Len() {
// 	fmt.Println("Cat", c.count)
// }

// func (d *dog) Len() {
// 	fmt.Println("Dog", d.count)
// }

// func (d *dog) Sound() {
// 	for d.count > 0 {
// 		fmt.Println("Вав!")
// 		d.count--
// 	}
// }

// func main() {
// 	var c, d animal = &cat{2}, &dog{3}
// 	c.Sound()
// 	d.Sound()
// 	c.Len()
// 	d.Len()
// }

// type greeter interface {
// 	greet(string) string
// }

// type russian struct{}

// type american struct{}

// func (r russian) greet(s string) string {
// 	return fmt.Sprintf("Привет, %s!", s)
// }
// func (a american) greet(s string) string {
// 	return fmt.Sprintf("Hello, %s!", s)
// }

// func SayHello(g greeter, s string) {
// 	fmt.Println(g.greet(s))
// }

// func main() {
// 	// var r, a greeter = russian{}, american{}
// 	SayHello(russian{}, "Петя")
// 	SayHello(american{}, "Alex")

// }

type Shape interface {
	Area() float32
}

type Square struct {
	iFloat32 float32
}

func (s Square) Area() float32 {
	return s.i * s.i
}

type Circle struct {
	i float32
}

func (c Circle) Area() float32 {
	return c.i * c.i * float32(math.Pi)
}

func main() {
	q := Square{3.12}
	c := Circle{242.12}
	Printer(q)
	Printer(c)

}
func Printer(s Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", s.Area())
}
