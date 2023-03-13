package main

import "fmt"

func main() {
	N := 0
	k := 0
	fmt.Scanf("%d %d", &N, &k)
	if N == 1 {
		fmt.Println(N)
		return
	}
	arr := make([]int, N)
	arr[0], arr[1] = 1, 1
	for m := 2; m < len(arr); m++ {
		maxlen := m
		if m > k {
			maxlen = k
		}
		sum := 0
		for ; maxlen >= 1; maxlen-- {
			sum += arr[m-maxlen]
		}
		arr[m] = sum
	}
	fmt.Println(arr[N-1])
}
