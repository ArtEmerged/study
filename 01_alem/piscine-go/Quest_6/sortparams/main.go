package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	test := os.Args[1:]
	// var R []rune
	// for _, n := range test {
	// 	R = append(R, rune(n[0]))
	// }

	Sort(test)
	for _, n := range test {
		for _, i := range n {
			z01.PrintRune(rune(i))
		}
		z01.PrintRune('\n')
	}
}

func Sort(R []string) {
	for n := 0; n < len(R); n++ {
		for k := 0; k < len(R)-1; k++ {
			if R[k][0] > R[k+1][0] {
				R[k], R[k+1] = R[k+1], R[k]
			}
		}
	}
}
