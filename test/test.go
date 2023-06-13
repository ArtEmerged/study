package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	printalphabet()
	printreversealphabet()
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
	PrintComb()
}

func printalphabet() {
	for a := 'a'; a <= 'z'; a++ {
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}
func printreversealphabet() {
	for a := 'z'; a >= 'a'; a-- {
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}

func IsNegative(nb int) {
	if nb < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func PrintComb() {
	for a := '0'; a < '7'; a++ {
		for b := '1'; a < '8'; b++ {
			for c := '2'; c <= '9'; c++ {
				if a < b && b < c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if a == '7' {
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
