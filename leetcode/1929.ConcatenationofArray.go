package leet

func getConcatenation(nums []int) []int {
	ans := make([]int, 0, len(nums)*2)
	ans = append(nums, nums...)
	return ans
}
