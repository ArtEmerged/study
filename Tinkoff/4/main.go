package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	max_diff := make([]int, k)
	for i := 0; i < n; i++ {
		var digit int
		fmt.Scan(&digit)
		rank := 1
		for digit > 0 {
			diff := digit % 10
			max := (9 - diff) * rank
			rank *= 10
			digit /= 10
			min := max_diff[0]
			ind := 0
			for j := 0; j < k; j++ {
				if min > max_diff[j] {
					min = max_diff[j]
					ind = j
				}
			}
			if max > min {
				max_diff[ind] = max
			}
		}
	}
	sum := 0
	for _, v := range max_diff {
		sum += v
	}
	fmt.Println(sum)
}
