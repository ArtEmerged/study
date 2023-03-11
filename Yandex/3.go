package main

import (
	"fmt"
	"sort"
)

func main() {
	var Sticks, Collector int
	fmt.Scanf("%d", &Sticks)
	NumberSt := make([]int, 0, Sticks)
	key := map[int]int{}
	for f := 0; f < Sticks; f++ {
		buf := 0
		fmt.Scanf("%d", &buf)
		key[buf]++
		if key[buf] == 1 {
			NumberSt = append(NumberSt, buf)
		}
	}
	fmt.Scanf("%d", &Collector)
	StickersColl := make([]int, Collector)
	for i := range StickersColl {
		fmt.Scanf("%d", &StickersColl[i])
	}
	sort.Ints(NumberSt)
	// sortInt(NumberSt)
	for _, number := range StickersColl {
		fmt.Println(Search(NumberSt, number))
	}
}

// func sortInt(arr []int) []int {
// 	if len(arr) > 1 {
// 		Left := 0
// 		Right := len(arr) - 1
// 		middle := arr[(Right+Left)/2]
// 		for Left < Right {
// 			for arr[Left] < middle {
// 				Left++
// 			}
// 			for arr[Right] > middle {
// 				Right--
// 			}
// 			if Left <= Right {
// 				arr[Left], arr[Right] = arr[Right], arr[Left]
// 				Left++
// 				Right--
// 			}
// 		}
// 		sortInt(arr[:Right+1])
// 		sortInt(arr[Left:])
// 	}
// 	return arr
// }

func Search(arr []int, find int) int {
	if find != 0 {
		Left := 0
		Right := len(arr) - 1
		answer := 0
		for Left <= Right {
			middle := (Right + Left) / 2
			if find == arr[middle] {
				return middle
			}
			if find > arr[middle] {
				Left = middle + 1
			} else {
				Right = middle - 1
			}
			answer = middle
		}
		if find > arr[answer] {
			return answer + 1
		}
		return answer
	}
	return 0
}
