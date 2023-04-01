package leet

func findKthPositive(arr []int, k int) int {
	var ans, i int
	for k > 0 {
		ans++
		if i < len(arr) && arr[i] == ans {
			i++
		} else {
			k--
		}
	}
	return ans
}
