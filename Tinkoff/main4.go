//#4
package main

import (
	"fmt"
	"sort"
)

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		var v int
		fmt.Scan(&v)
		arr[i] = v
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	money := []int{}
	for j := 0; j < m; j++ {
		count := 0
		for n >= arr[j] && count < 2 {
			money = append(money, arr[j])
			n = n - arr[j]
			count++
		}
	}
	sort.Ints(money)
	if n == 0 {
		fmt.Println(len(money))
		for _, v := range money {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	} else {
		fmt.Println(-1)
	}
}
