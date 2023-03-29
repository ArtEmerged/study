package leet

func sumZero(n int) []int {
	answer := make([]int, n)
	for i := 0; i < (n - n%2); i += 2 {
		answer[i] = n + i
		answer[i+1] = -(n + i)
	}
	return answer
}
