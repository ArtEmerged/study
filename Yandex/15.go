// package main

// import "fmt"

// // type LIFO struct {
// // 	Num   int
// // 	Index int
// // 	Next  *LIFO
// // }

// // type List struct {
// // 	Head *LIFO
// // }

// func main() {
// 	lenght := 10
// 	// fmt.Scanln(&lenght)
// 	// answer := make([]int, lenght)
// 	answer2 := []int{1, 2, 3, 2, 1, 4, 2, 5, 3, 1}
// 	answer := make([]int, lenght)
// 	arr := make([]int, 0, lenght)
// 	index := 0
// 	for i, n := range answer2 {
// 		// n := 0
// 		// fmt.Scan(&n)
// 		for len(arr) != 0 && n < arr[len(arr)-1] {
// 			index--
// 			answer[index] = i
// 			arr = arr[:len(arr)-1]
// 		}
// 		arr = append(arr, n)
// 		index++
// 	}

// 	for _, v := range answer {
// 		if v == 0 {
// 			fmt.Print(-1)
// 		} else {
// 			fmt.Print(v)
// 		}
// 		fmt.Print(" ")
// 	}
// }

package main

import (
	"fmt"
	"time"
)

type LIFO struct {
	Num   int
	Index int
	Next  *LIFO
}

func main() {
	t0 := time.Now()
	lenght := 0
	fmt.Scanln(&lenght)
	arr := make([]int, lenght)
	l := &LIFO{}
	for i := range arr {
		n := 0
		fmt.Scan(&n)
		for l != nil && n < l.Num {
			arr[l.Index] = i
			l = l.Next
		}
		push := &LIFO{Num: n, Index: i, Next: l}
		l = push
	}
	for _, v := range arr {
		if v == 0 {
			fmt.Print(-1, " ")
		} else {
			fmt.Print(v, " ")
		}
	}
	t1 := time.Now()
	fmt.Printf("\nElapsed time: %v\n", t1.Sub(t0))
}
