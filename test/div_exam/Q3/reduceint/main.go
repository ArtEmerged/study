package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2, 8}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	n := a[0]
	for i := 1; i < len(a); i++ {
		n = f(n, a[i])
	}
	for _, val := range strconv.Itoa(n) {
		z01.PrintRune(val)
	}
	z01.PrintRune(10)
}




// func ReduceInt(a []int, f func(int, int) int) {
// 	res := f(a[0], a[1])
// 	resRune := ""
// 	if res == 0 {
// 		z01.PrintRune('0')
// 		return
// 	}
// 	if res < 0 {
// 		z01.PrintRune('-')
// 		res = res * -1
// 	}
// 	for res != 0 {
// 		resRune = string(res%10+48) + resRune
// 		res = res / 10
// 	}
// 	for _, n := range resRune {
// 		z01.PrintRune(n)
// 	}
// 	z01.PrintRune(10)
// }

