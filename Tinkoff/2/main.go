package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	count := 0
	f := 1
	for f < N {
		count++
		f *= 2
	}
	fmt.Println(count)
}
