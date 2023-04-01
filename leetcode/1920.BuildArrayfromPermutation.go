package leet

func buildArray(nums []int) []int {
	n := len(nums)
	for i := range nums {
		nums[i] += n * (nums[nums[i]] % n)
	}
	for i := range nums {
		nums[i] /= n
	}
	return nums
}
