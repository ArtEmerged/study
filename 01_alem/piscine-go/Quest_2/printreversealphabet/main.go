package main

import "github.com/01-edu/z01"

func main() {
	for v := 'z'; v >= 'a'; v-- {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
