package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type LIFO struct {
	Num   int
	Index int
	Next  *LIFO
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)
	lenght := 0
	fmt.Fscan(in, &lenght)
	arr := make([]int, lenght)
	l := &LIFO{}
	for i := range arr {
		n := 0
		fmt.Fscan(in, &n)
		for l != nil && n < l.Num {
			arr[l.Index] = i
			l = l.Next
		}
		push := &LIFO{Num: n, Index: i, Next: l}
		l = push
	}
	// file, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// out := bufio.NewWriter(file)
	// for _, v := range arr {
	// 	fmt.Fprintf(out, "%d ", v)
	// }

	for _, v := range arr {
		if v == 0 {
			fmt.Print(-1, " ")
		} else {
			fmt.Print(v, " ")
		}
	}

}
