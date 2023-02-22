package main

import "fmt"

func main() {

	
	fmt.Println(recursion(10))
}

func recursion(i int) int {
	if i == 0 {
		return 0
	}
	fmt.Println(i)
	recursion(i - 1)
	return 0
}
