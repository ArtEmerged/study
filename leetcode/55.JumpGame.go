package leet

func canJump(nums []int) bool {
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= tempIndex {
			tempIndex = i
		}
	}
	if tempIndex == 0 {
		return true
	}
	return false
}
