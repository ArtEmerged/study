package main

import "fmt"

func main() {
	sortedArr := []int{4, 7, 9, 12, 65, 78, 81, 84}
	fmt.Println(BinaryRecursion(sortedArr, -65))
}

//////////////////////////////////////
///////RECURSION SEARCH//////////////
////////////////////////////////////
func BinaryRecursion(arr []int, find int) int {
	return search(arr, find, 0, len(arr)-1)
}

func search(arr []int, find, left, right int) int {
	centre := (right + left) / 2
	middle := arr[centre]
	if middle == find && right-left == 0 {
		return centre
	} else if right-left != 0 {
		if find > middle {
			left = centre + 1
			return search(arr, find, left, right)
		} else {
			right = centre
			return search(arr, find, left, right)
		}
	}
	return 0
}

//////////////////////////////////////
///////ITERATIVE SEARCH//////////////
////////////////////////////////////

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
