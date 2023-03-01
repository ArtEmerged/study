package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					if (a*10+b != c*10+d) && (a*10+b < c*10+d) {
						z01.PrintRune(rune(a + 48))
						z01.PrintRune(rune(b + 48))
						z01.PrintRune(' ')
						z01.PrintRune(rune(c + 48))
						z01.PrintRune(rune(d + 48))
						if a == 9 && b == 8 && c == 9 && d == 9 {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
