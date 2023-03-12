package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	Args := os.Args[1:]
	Upp := false
	if Args[0] == "--upper" {
		Args = Args[1:]
		Upp = true
	}
	R := SrtInInt(Args)

	Output(R, Upp)
}

func SrtInInt(str []string) []int {
	R := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		for n := 0; n < len(str[i]); n++ {
			R[i] = R[i]*10 + int(str[i][n]-48)
		}
	}

	return R
}

func Output(R []int, upp bool) {
	Aa := 96
	if upp {
		Aa = 64
	}
	for _, n := range R {
		if n < 0 || n > 26 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune(rune(n + Aa))
		}
	}
	z01.PrintRune('\n')
}
