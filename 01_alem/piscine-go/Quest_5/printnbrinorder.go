package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		b := Move(n)
		Array := Sort(b)

		for _, i := range Array {
			z01.PrintRune(rune(i + 48))
		}
	}
}

func Move(n int) []int {
	var b []int
	for n > 0 {
		b = append(b, n%10)
		n /= 10
	}
	return b
}

func Sort(b []int) []int {
	for i := 0; i < len(b); i++ {
		for z := 0; z < len(b)-1-i; z++ {
			if b[z] > b[z+1] {
				b[z], b[z+1] = b[z+1], b[z]
			}
		}
	}
	return b
}
