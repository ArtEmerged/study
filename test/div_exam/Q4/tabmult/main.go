package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		res := ""
		n, _ := strconv.Atoi(os.Args[1])
		for i := 1; i < 10; i++ {
			res += strconv.Itoa(i) + " + " + strconv.Itoa(n) + " = " + strconv.Itoa(i*n) + "\n"
		}
		for _, n := range res {
			z01.PrintRune(n)
		}
	}
}
