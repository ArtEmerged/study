package leet

import "sort"

func sortedSquares(nums []int) []int {

	for i, n := range nums {
		nums[i] = n * n
	}
	sort.SliceStable(nums, func(i, j int) bool { return nums[j] > nums[i] })
	return nums
}
