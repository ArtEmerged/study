package main

import "fmt"

func main() {

	array := []int{4, 2, 5, 3, 6, 3, 66, 32, 66, 454, 22}

	fmt.Println(max(array))

}

func max(array []int) int {
	if len(array) == 1 {
		return array[0]
	} else if len(array) == 0 {
		return 0
	} else if len(array) == 2 {
		if array[0] > array[1] {
			return array[0]
		}
		return array[1]
	}

	maxM := max(array[1:])
	if maxM > array[0] {
		return maxM
	}
	return array[0]
}
