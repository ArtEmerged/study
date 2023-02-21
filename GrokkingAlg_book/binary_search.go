package main

import "fmt"

func main() {
	hundred := make([]int, 100)
	for n := 1; n <= len(hundred); n++ {
		hundred[n-1] = n
	}
	fmt.Println(hundred)
	fmt.Println(BinaryIterat(hundred, 3))
}

func BinaryIterat(arr []int, find int) int {
	min := 0
	max := len(arr) - 1
	for min <= max {
		middle := (min + max) / 2
		if arr[middle] == find && max-min == 0 {
			return middle
		}
		if find > arr[middle] {
			min = middle + 1
		} else {
			max = middle
		}
	}
	return 0
}
