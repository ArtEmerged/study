package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[0]
	// A := []string{"link", "Hello"}
	// for ind, o := range a[0] {
	// 	if o != '.' && o != '/' {
	// 		z01.PrintRune(rune(a[0][ind]))
	// 	}
	// }
	// for ind := 2; ind < len(a[0]); ind++ {
	// 	z01.PrintRune(rune(a[0][ind]))
	// }
	for _, v := range a[2:] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
