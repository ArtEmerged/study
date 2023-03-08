package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)
	var v string
	l := &list{}
	for {
		_, err = fmt.Fscan(in, &v)
		if err == io.EOF {
			break
		}
		if v == " " {
			continue
		}
		if v >= "0" && v <= "9" {
			i := &LIFO{Num: int(v[0] - 48), Next: l.Head}
			l.Head = i
		} else {
			switch v {
			case "+":
				Val := l.Head.Next.Num + l.Head.Num
				l.Head = l.Head.Next
				l.Head.Num = Val
			case "-":
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
