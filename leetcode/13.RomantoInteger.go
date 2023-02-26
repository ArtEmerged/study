package leet

func romanToInt(s string) int {
	key := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	answer := key[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if key[s[i]] < key[s[i+1]] {
			answer -= key[s[i]]
		} else {
			answer += key[s[i]]
		}
	}
	return answer

}
