package leet

func smallerNumbersThanCurrent(nums []int) []int {
	answer := make([]int, len(nums))
	for i, v := range nums {
		for _, n := range nums {
			if v > n {
				answer[i]++
			}
		}
	}
	return answer
}

