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

	fmt.Println("\nPrintComb")
	PrintComb()

	fmt.Println("\nPrintComb2")
	PrintComb2()

	fmt.Println("\nPrintNbr")
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')

	fmt.Println("\nPointOne")
	n := 0
	PointOne(&n)
	fmt.Println(n)

	fmt.Println("\nUltimatePointOne")
	a := 0
	b := &a
	k := &b
	UltimatePointOne(&k)
	fmt.Println(a)

	fmt.Println("\nDivMod")
	h := 13
	j := 2
	var div int
	var mod int
	DivMod(h, j, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)

	fmt.Println("\nUltimateDivMod")
	aa := 13
	bb := 2
	UltimateDivMod(&aa, &bb)
	fmt.Println(aa)
	fmt.Println(bb)

	fmt.Println("\nPrintStr")
	PrintStr("Hello World!")

	fmt.Println("\nStrLen")
	l := StrLen("Hello World!")
	fmt.Println(l)

	fmt.Println("\nSwap")
	aaa := 0
	bbb := 1
	Swap(&aaa, &bbb)
	fmt.Println(aaa)
	fmt.Println(bbb)

	fmt.Println("\nStrRev")
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)

	fmt.Println("\nBasicAtoi")
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))

	fmt.Println("\nBasicAtoi2")
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))

	fmt.Println("\nAtoi")
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
	fmt.Println(Atoi("-"))

	fmt.Println("\nIterativeFactorial")
	arg := 4
	fmt.Println(IterativeFactorial(arg))

	fmt.Println("\nRecursiveFactorial")
	fmt.Println(RecursiveFactorial(arg))

	fmt.Println("\n IterativePower")
	fmt.Println(IterativePower(4, 3))

	fmt.Println("\n RecursivePower")
	fmt.Println(RecursivePower(4, 3))

	fmt.Println("\n Fibonacci")
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}

// QUEST 02
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
	for i := '0'; i <= '7'; i++ {
		for b := '1'; b <= '8'; b++ {
			for c := '2'; c <= '9'; c++ {
				if i < b && b < c {
					z01.PrintRune(i)
					z01.PrintRune(b)
					z01.PrintRune(c)
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

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					if (a*10 + b) < (c*10 + d) {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(c)
						z01.PrintRune(d)
						if a == '9' && b == '8' && c == '9' && d == '9' {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}

func PrintNbr(n int) {
	var res []rune
	var last int
	if n == 0 {
		z01.PrintRune('0')
		return
	} else if n < 0 {
		res = append(res, rune(n%10*-1+48))
		last = n / 10 * -1
		z01.PrintRune('-')
	} else {
		last = n
	}
	for last != 0 {
		res = append(res, rune(last%10+48))
		last /= 10
	}
	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(res[i])
	}
}

// QUEST 03
func PointOne(n *int) {
	*n = 1
}

func UltimatePointOne(n ***int) {
	***n = 1
}

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}

func UltimateDivMod(aa *int, bb *int) {
	*aa, *bb = *aa / *bb, *aa%*bb
}

func PrintStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}

func StrLen(s string) int {
	return len(s)
}

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

func StrRev(s string) string {
	b := make([]rune, len(s))
	rev := len(b) - 1
	for n := 0; n < len(s); n++ {
		b[rev] = rune(s[n])
		rev--
	}
	return string(b)
}

func BasicAtoi(s string) int {
	var res int
	for _, i := range s {
		res = res*10 + int(i-48)
	}
	return res
}

func BasicAtoi2(s string) int {
	var res int
	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		}
		res = res*10 + int(i-48)
	}
	return res
}

func Atoi(s string) int {
	var res int
	sign := 1
	for n, i := range s {
		if n == 0 {
			if i == '-' {
				sign = -1
				continue
			} else if i == '+' {
				continue
			}
		}
		if i < '0' || i > '9' {
			return 0
		}
		res = res*10 + int(i-48)
	}
	return res * sign
}

func IterativeFactorial(nb int) int {
	res := 1
	for i := 1; i <= nb; i++ {
		res *= i
	}
	return res
}

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	} else if nb < 0 || nb > 20 {
		return 0
	}
	return nb * RecursiveFactorial(nb-1)
}

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	}
	res := nb
	for i := 1; i < power; i++ {
		res *= nb
	}
	return res
}

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	}
	return nb * RecursivePower(nb, power-1)
}

func Fibonacci(index int) int {
	var result int
	res1 := 0
	res2 := 1
	if index < 0 {
		return -1
	}
	for i := 0; i < index; i++ {
		result = res2 + res1
		res1 = res2
		res2 = result
	}
	return result
}
