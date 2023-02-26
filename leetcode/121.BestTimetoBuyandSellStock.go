package leet

import "math"

func maxProfit(prices []int) int {
	min, answer := math.MaxInt64, 0
	for _, cost := range prices {
		if min > cost {
			min = cost
		} else if cost-min > answer {
			answer = cost - min
		}
	}
	return answer
}
