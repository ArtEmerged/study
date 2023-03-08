package main

import "fmt"

type LIFO struct {
	Num   int
	Index int
	Next  *LIFO
}

type List struct {
	Head *LIFO
}

func main() {
	lenght := 0
	fmt.Scanln(&lenght)
	arr := make([]int, lenght)
	l := &List{}
	for i := range arr {
		n := 0
		fmt.Scan(&n)
		for l.Head != nil && n < l.Head.Num {
			arr[l.Head.Index] = i
			l.Head = l.Head.Next
		}
		add := &LIFO{Num: n, Index: i, Next: l.Head}
		l.Head = add
	}
	
	for _, v := range arr {
		if v == 0 {
			fmt.Print(-1)
		} else {
			fmt.Print(v)
		}
		fmt.Print(" ")
	}
}
