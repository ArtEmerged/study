package main

import (
	"fmt"
)

type Stek struct {
	num  int
	next *Stek
}

type List struct {
	Head *Stek
	Tail *Stek
}

func main() {
	exit := true
	l := &List{}
	size := 0
	for exit {
		var step string
		var number int
		fmt.Scanln(&step, &number)
		switch step {
		case "push": //+
			size++
			if l.Head != nil {
				i := &Stek{num: number}
				l.Tail.next = i
				l.Tail = l.Tail.next
			} else {
				i := &Stek{num: number}
				l.Head = i
				l.Tail = i
			}
			fmt.Println("ok")
		case "pop": //+
			if l.Head != nil {
				fmt.Println(l.Head.num)
				l.Head = l.Head.next
				size--
			} else {
				fmt.Println("error")
			}
		case "front": //+
			if l.Head != nil {
				fmt.Println(l.Head.num)
			} else {
				fmt.Println("error")
			}
		case "size":
			fmt.Println(size)
		case "clear": //+
			for l.Head != nil {
				l.Head = l.Head.next
			}
			size = 0
			fmt.Println("ok")
		default: //+
			fmt.Println("bye")
			exit = false
		}
	}
}
