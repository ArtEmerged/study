package main

import "fmt"

func main() {

	array := []int{1, 3, 4, 9, 3}
	fmt.Println(sum(array))
}

func sum(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return array[0] + sum(array[1:])

}
