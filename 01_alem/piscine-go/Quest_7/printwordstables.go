package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, i := range a {
		for _, n := range i {
			z01.PrintRune(n)
		}
		z01.PrintRune('\n')
	}
}
