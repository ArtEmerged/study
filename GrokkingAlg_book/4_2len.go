package main

import "fmt"

func main() {

	array := []int{4, 2, 5, 3, 6, 3, 66, 32, 66, 454, 22}

	fmt.Println(countLen(array))

}

func countLen(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return 1 + countLen(array[1:])
}
