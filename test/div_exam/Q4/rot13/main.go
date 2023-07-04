package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, n := range os.Args[1] {
			if n >= 'a' && n <= 'm' || n >= 'A' && n <= 'M' {
				z01.PrintRune(n + 13)
			} else if n >= 'n' && n <= 'z' || n >= 'N' && n <= 'Z' {
				z01.PrintRune(n - 13)
			} else {
				z01.PrintRune(n)
			}
		}
	}
	z01.PrintRune(10)
}
