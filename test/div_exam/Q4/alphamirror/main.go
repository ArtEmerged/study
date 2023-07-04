package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, n := range os.Args[1] {
			if n >= 'a' && n <= 'z' {
				n = 'z' - (n - 'a')
			} else if n >= 'A' && n <= 'Z' {
				n = 'Z' - (n - 'A')
			}
			z01.PrintRune(n)
		}
	}
	z01.PrintRune(10)
}
