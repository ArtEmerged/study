package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for f := '1'; f <= '8'; f++ {
			for k := '2'; k <= '9'; k++ {
				if f < k && i < f {
					z01.PrintRune(i)
					z01.PrintRune(f)
					z01.PrintRune(k)
					if i == '7' {
						break
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
