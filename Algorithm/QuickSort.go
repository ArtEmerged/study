package main

import "fmt"

func main() {

	number := []int{99994, 1, 9, 0, 7, 8, 1, 54, 34, 3, 1, 9, 88, 15, 90, 18, 22, 22, 22, 15, 4, 1, 87, 3, 454}

	fmt.Println(QuickSort(number))

}

func QuickSort(arr []int) []int {
	if len(arr) > 1 {
		left, right := 0, len(arr)-1
		middle := arr[(left+right)/2]
		for left < right {
			for arr[left] < middle {
				left++
			}
			for arr[right] > middle {
				right--
			}
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
		QuickSort(arr[:right+1])
		QuickSort(arr[left:])
	}
	return arr

}
