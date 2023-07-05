package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		word := os.Args[1]
		letterUp := "AEIOU"
		letterLow := "aeiou"
		for i := range word {
			for l := range letterUp {
				if word[i] == letterLow[l] || word[i] == letterUp[l] {
					word = word[i:] + word[:i] + "ay"
					for _, w := range word {
						z01.PrintRune(w)
					}
					z01.PrintRune(10)
					return
				}
			}
		}
		for _, w := range "No vowels" {
			z01.PrintRune(w)
		}
		z01.PrintRune(10)

	}
}
