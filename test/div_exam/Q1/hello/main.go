package main

import "github.com/01-edu/z01"

func main() {
	hello := "Hello World!"
	for _, n := range hello {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}
