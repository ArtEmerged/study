// package main

// import (
// 	"fmt"
// 	"sort"
// )

// var (
// 	n int
// 	m int
// )

// func main() {
// 	fmt.Scanf("%d %d", &n, &m)
// 	max := m
// 	arr := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&arr[i])
// 	}
// 	for i := 0; i < n; i++ {
// 		lMax := arr[i] - 1
// 		rMax := max - arr[i]
// 		max = fMax(lMax, rMax)
// 	}
// 	fmt.Println(max)
// }

// func fMax(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	fmt.Scan(&n, &m)

// 	prices := make([]int, m)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&prices[i])
// 	}

// 	sort.Ints(prices)

// 	maxRemaining := getMaxRemaining(n, m, prices)
// 	fmt.Println(maxRemaining)
// }

// func getMaxRemaining(n, m int, prices []int) int {
// 	maxRemaining := 0
// 	for i := 0; i < n; i++ {
// 		if prices[i] <= maxRemaining {
// 			maxRemaining++
// 		}
// 	}
// 	return m - maxRemaining
// }

package main

import (
	"fmt"
	"sort"
)

func maxCoinsAfterShopping(n, m int, prices []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(prices))) // Сортируем цены по убыванию

	remainingCoins := m
	i := 0

	for i < n && remainingCoins >= prices[i] {
		remainingCoins -= prices[i]
		i++
	}

	return remainingCoins
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&prices[i])
	}

	result := maxCoinsAfterShopping(n, m, prices)
	fmt.Println(result)
}
