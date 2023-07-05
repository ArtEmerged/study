package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			for _, v := range "00000000" {
				z01.PrintRune(v)
			}
			z01.PrintRune(10)
			return
		}
		res := ""
		for len(res) != 8 {
			res = string(n%2+48) + res
			n /= 2
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	}
}
