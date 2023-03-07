package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	b := len(base)
	if CheckValid(base, b) {
		NotValid()
		return
	}
	var answer []rune
	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}
	for nbr > 0 {
		answer = append(answer, rune(base[nbr%b]))
		nbr /= b
	}
	for i := len(answer) - 1; i >= 0; i-- {
		z01.PrintRune(answer[i])
	}
}
func NotValid() {
	z01.PrintRune('N')
	z01.PrintRune('V')
}
func CheckValid(base string, ln int) bool {
	if ln < 2 {
		return true
	}
	for i := 0; i < ln-1; i++ {
		for n := i + 1; n < ln; n++ {
			if base[i] == base[n] {
				return true
			}
		}
	}
	for _, val := range base {
		if val == '+' || val == '-' {
			return true
		}
	}
	return false
}
