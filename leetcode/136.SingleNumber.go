package leet

import "sort"

func singleNumber(nums []int) int {
	lenght := len(nums)
	if lenght == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	for i := 0; i < lenght-1; i = i + 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[lenght-1]
}
