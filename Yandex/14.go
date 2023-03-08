package main

import (
	"fmt"
	"math"
)

type LIFO struct {
	Num  int
	Next *LIFO
}

type List struct {
	Head *LIFO
}

func main() {
	min := math.MaxInt64
	var lenght int
	fmt.Scanln(&lenght)
	N := make([]int, lenght)
	l := &List{}
	for i := range N {
		var num int
		fmt.Scan(&num)
		N[i] = num
		if num < min {
			min = num
		}
	}
	for _, v := range N {
		i := &LIFO{Num: v, Next: l.Head}
		l.Head = i
		if l.Head.Num == min {
			l.Head = l.Head.Next
			min++
		}
	}
	for l.Head != nil {
		if l.Head.Num == min {
			l.Head = l.Head.Next
			min++
		} else {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
