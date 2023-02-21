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

func BinaryIterat(hundred []int, m int) int {
	min := 0
	max := len(hundred) - 1
	for min <= max {
		midddle := (min + max) / 2

		if hundred[midddle] == m {
			return midddle
		}
		if hundred[midddle] < m {
			min = midddle + 1
		} else {
			max = midddle - 1
		}

	}
	return 0

}
