package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		check := map[rune]int{}
		for _, v := range os.Args[1:] {
			for _, r := range v {
				if _, ch := check[r]; !ch {
					check[r]++
					z01.PrintRune(r)
				}
			}
		}
	}
	z01.PrintRune(10)
}
