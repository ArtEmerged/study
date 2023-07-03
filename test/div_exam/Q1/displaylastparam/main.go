package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) != 1 {

		last := len(arg) - 1
		for _, i := range arg[last] {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
