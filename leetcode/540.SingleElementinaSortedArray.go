package leet

func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if (mid-1 < 0 || nums[mid-1] != nums[mid]) && (mid+1 == len(nums) || nums[mid+1] != nums[mid]) {
			return nums[mid]
		}
		leftSize := mid
		if nums[mid-1] == nums[mid] {
			leftSize = mid - 1
		}
		if leftSize%2 == 1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
