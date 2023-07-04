package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		a1, a2 := rune(os.Args[2][0]), rune(os.Args[3][0])
		for _, n := range os.Args[1] {
			if n == a1 {
				n = a2
			}
			z01.PrintRune(n)
		}
		z01.PrintRune(10)
	}
}
