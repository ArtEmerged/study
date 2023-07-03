package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		return
	}
	for _, i := range arg[1] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
