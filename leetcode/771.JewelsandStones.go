package leet

func numJewelsInStones(jewels string, stones string) int {
	answer := 0
	for _, v := range jewels {
		for _, m := range stones {
			if v == m {
				answer++
			}
		}
	}
	return answer
}
