package main

import "fmt"

func main() {

	array := []int{4, 2, 5, 3, 6, 3, 66, 32, 66, 454, 22}


	fmt.Println(findMax(array))
	fmt.Println(findMax2(array))

}

func findMax(array []int) int {
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

	maxM := findMax(array[1:])
	if maxM > array[0] {
		return maxM
	}
	return array[0]
}


//version 2
func findMax2(x []int) int {
	if len(x) == 0 {
		return 0
	}
	if max := findMax2(x[1:]); max > x[0] {
		return max
	}
	return x[0]
}
