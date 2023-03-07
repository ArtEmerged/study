package main

import (
	"fmt"
)

type Stek struct {
	num  int
	size int
	next *Stek
}

type List struct {
	Head *Stek
}

func main() {
	exit := true
	l := &List{}
	for exit {
		var step string
		var number int
		fmt.Scanln(&step, &number)
		switch step {
		case "push": //+
			if l.Head != nil {
				i := &Stek{num: number, size: l.Head.size, next: l.Head}
				l.Head = i
				l.Head.size++
			} else {
				i := &Stek{num: number, next: l.Head}
				l.Head = i
				l.Head.size++
			}
			fmt.Println("ok")
		case "pop": //+
			if l.Head != nil {
				fmt.Println(l.Head.num)
				l.Head = l.Head.next
			} else {
				fmt.Println("error")
			}
		case "back": //+
			if l.Head != nil {
				fmt.Println(l.Head.num)
			} else {
				fmt.Println("error")
			}
		case "size":
			if l.Head != nil {
				fmt.Println(l.Head.size)
			} else {
				fmt.Println(0)
			}
		case "clear": //+
			for l.Head != nil {
				l.Head = l.Head.next
			}
			fmt.Println("ok")
		default: //+
			fmt.Println("bye")
			exit = false
		}
	}
}
