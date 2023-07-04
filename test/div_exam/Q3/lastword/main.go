package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		result := ""
		for i := 0; i < len(arg); i++ {
			if arg[i] != ' ' {
				result += string(arg[i])
			} else if i+1 < len(arg) && (arg[i+1] != ' ') {
				result = ""
			}
		}
		for _, n := range result {
			z01.PrintRune(n)
		}
		z01.PrintRune('\n')
	}
}
