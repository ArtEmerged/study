package main

import "github.com/01-edu/z01"

func main() {
	for v := 'a'; v <= 'z'; v++ {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
