package main

import "github.com/01-edu/z01"

func main() {
	for a := 'z'; a >= 'a'; a-- {
		if a%2 == 1 {
			z01.PrintRune(a - 32)
		} else {
			z01.PrintRune(a)
		}
	}
	z01.PrintRune('\n')
}
