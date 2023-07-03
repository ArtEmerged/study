package main

import "github.com/01-edu/z01"

func main() {
	for a := 'a'; a <= 'z'; a++ {
		if a%2 == 0 {
			z01.PrintRune(a - 32)
		} else {
			z01.PrintRune(a)
		}
	}
	z01.PrintRune('\n')
}
