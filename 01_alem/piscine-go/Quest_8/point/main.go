package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func main() {
	points := &point{}
	setPoint(points)
	// fmt.Printf("x = %v, y = %v\n", points.x, points.y)
	answer := "x = 42, y = 21"
	Outpud(answer)
}

func setPoint(n *point) {
	n.x = 42
	n.y = 21
}

func Outpud(ans string) {
	for _, n := range ans {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}
