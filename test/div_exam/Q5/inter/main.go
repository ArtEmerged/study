package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		check := map[rune]bool{}
		for _, r1 := range os.Args[1] {
			for _, r2 := range os.Args[2] {
				if r1 == r2 {
					if _, val := check[r1]; !val {
						check[r1] = true
						z01.PrintRune(r1)
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}
