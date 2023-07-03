package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) < 10 {
		z01.PrintRune(rune(len(os.Args[1:])) + 48)
	} else {
		a := len(os.Args[1:])
		z01.PrintRune(rune(a/10 + 48))
		z01.PrintRune(rune(a%10 + 48))
	}
	z01.PrintRune('\n')
}
