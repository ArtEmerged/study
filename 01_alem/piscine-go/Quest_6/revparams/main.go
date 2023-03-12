package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	A := os.Args
	l := len(A)
	for i := l - 1; i > 0; i-- {
		for j := 0; j < len(A[i]); j++ {
			z01.PrintRune(rune(A[i][j]))
		}
		z01.PrintRune('\n')
	}
}
