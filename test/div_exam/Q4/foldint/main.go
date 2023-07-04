package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func Add(acc, cur int) int {
	return acc + cur
}

func Mul(acc, cur int) int {
	return acc * cur
}

func Sub(acc, cur int) int {
	return acc - cur
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for i := 0; i < len(a); i++ {
		n = f(n, a[i])
	}
	for _, val := range Itoi(n) {
		z01.PrintRune(val)
	}
	z01.PrintRune(10)
}

func Itoi(n int) string {
	neg := ""
	if n < 0 {
		neg = "-"
		n = n * -1
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	res := ""
	for n != 0 {
		res = string(rune(n%10+48)) + res
		n = n / 10
	}
	return neg + res
}
