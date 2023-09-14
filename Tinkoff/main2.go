// #2
package main

import (
	"fmt"
	"sort"
)

func main() {
	var str string
	arr := []int{0, 0, 0, 0, 0, 0}
	fmt.Scan(&str)
	for _, v := range str {
		switch v {
		case 's':
			arr[0]++
		case 'h':
			arr[1]++
		case 'e':
			arr[2]++
		case 'r':
			arr[3]++
		case 'i':
			arr[4]++
		case 'f':
			arr[5]++
		}
	}
	arr[5] = arr[5] / 2
	sort.Ints(arr)
	fmt.Println(arr[0])
}
