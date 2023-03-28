package leet

func sortArrayByParity(nums []int) []int {
	any1 := make([]int, 0, len(nums)/2)
	any2 := make([]int, 0, len(nums)/2)
	for _, n := range nums {
		if n%2 == 1 {
			any2 = append(any2, n)
		} else {
			any1 = append(any1, n)
		}
	}
	any1 = append(any1, any2...)
	return any1
}
