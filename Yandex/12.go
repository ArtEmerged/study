package main

import (
	"fmt"
)

type Stek struct {
	bracket rune
	next    *Stek
}

type List struct {
	Head *Stek
}

func main() {
	l := &Stek{}
	var str string
	fmt.Scanln(&str)
	for _, v := range str {
		switch v {
		case '(':
			i := &Stek{bracket: v, next: l.Head}
			l.Head = i
		case '[':
			i := &Stek{bracket: v, next: l.Head}
			l.Head = i
		case '{':
			i := &Stek{bracket: v, next: l.Head}
			l.Head = i
		case ')':
			if l.Head == nil || l.Head.bracket != '(' {
				fmt.Println("no")
				return
			} else {
				l.Head = l.Head.next
			}
		case ']':
			if l.Head == nil || l.Head.bracket != '[' {
				fmt.Println("no")
				return
			} else {
				l.Head = l.Head.next
			}
		default:
			if l.Head == nil || l.Head.bracket != '{' {
				fmt.Println("no")
				return
			} else {
				l.Head = l.Head.next
			}
		}
	}
	if l.Head == nil {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
