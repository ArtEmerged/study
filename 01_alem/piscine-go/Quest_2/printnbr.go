package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var runs []rune

	Last := 0
	if n == 0 {
		z01.PrintRune('0')
		return
	} else if n > 0 {
		Last = n
	} else if n < 0 {
		runs = append(runs, rune(n%10*-1))
		Last = n / 10 * -1
		// n *= -1
		z01.PrintRune('-')
	}

	for Last > 0 {
		runs = append(runs, rune(Last%10))
		Last /= 10
	}
	for i := len(runs) - 1; i >= 0; i-- {
		z01.PrintRune(runs[i] + 48)
	}
}
