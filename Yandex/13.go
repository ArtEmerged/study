package main

import (
	"fmt"
	"os"
)

type LIFO struct {
	Num  int
	Next *LIFO
}

type list struct {
	Head *LIFO
}

func main() {
	line, _ := os.ReadFile("input.txt")
	line2 := string(line)
	l := &list{}
	for _, v := range line2 {
		if v == ' ' {
			continue
		}
		if v >= '0' && v <= '9' {
			i := &LIFO{Num: int(v - 48), Next: l.Head}
			l.Head = i
		} else {
			switch v {
			case '+':
				Val := l.Head.Next.Num + l.Head.Num
				l.Head = l.Head.Next
				l.Head.Num = Val
			case '-':
				Val := l.Head.Next.Num - l.Head.Num
				l.Head = l.Head.Next
				l.Head.Num = Val
			default:
				Val := l.Head.Next.Num * l.Head.Num
				l.Head = l.Head.Next
				l.Head.Num = Val
			}
		}
	}
	fmt.Println(l.Head.Num)
}
