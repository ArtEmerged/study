package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	Rune := []rune(s)

	for _, i := range Rune {
		z01.PrintRune(i)
	}
}
