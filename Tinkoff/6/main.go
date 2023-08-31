package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	steck := []int{}
	for i := 1; i <= n; i++ {
		growth := 0
		fmt.Scan(&growth)
		if growth%2 != i%2 {
			steck = append(steck, i)
		}
	}
	if len(steck) == 2 && steck[0]%2 != steck[1]%2 {
		fmt.Printf("%d %d\n", steck[0], steck[1])
		return
	}
	fmt.Println("-1 -1")
}
