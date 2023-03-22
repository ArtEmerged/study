package leet

func search(nums []int, target int) int {
	for i, num := range nums {
		if target == num {
			return i
		}
	}
	return -1
}
