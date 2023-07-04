package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	arr := [][]int{}
	if len(slice) == 0 {
		fmt.Println()
		return
	}
	if size <= 0 {
		fmt.Println()
		return
	}
	for len(slice) > size {
		arr = append(arr, slice[:size])
		slice = slice[size:]
	}
	arr = append(arr, slice)
	fmt.Println(arr)
}
