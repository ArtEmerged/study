package main

import "fmt"

func main() {
	var n, t, pos int
	fmt.Scanf("%d %d", &n, &t)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&pos)
	pos--
	last := len(arr) - 1
	if arr[pos]-arr[0] <= t {
		fmt.Println(arr[last] - arr[0])
	} else if arr[last]-arr[pos] <= t {
		fmt.Println(arr[last] - arr[0])
	} else if arr[pos]-arr[0] > arr[last]-arr[pos] {
		res := (arr[last] - arr[pos]) + (arr[last] - arr[0])
		fmt.Println(res)
	} else {
		res := (arr[pos] - arr[0]) + (arr[last] - arr[0])
		fmt.Println(res)
	}
}
