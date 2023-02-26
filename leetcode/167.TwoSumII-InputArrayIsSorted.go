package leet

func twoSum(numbers []int, target int) []int {
	answer := make([]int, 2)
	for n := 0; n < len(numbers)-1; n++ {

		for i := n + 1; i < len(numbers); i++ {
			if numbers[n]+numbers[i] == target {
				answer[0], answer[1] = n+1, i+1
				break

			}

		}

	}
	return answer
}