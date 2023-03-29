package leet

func defangIPaddr(address string) string {
	answer := ""
	for _, v := range address {
		if v == '.' {
			answer += "[.]"
		} else {
			answer += string(v)
		}
	}
	return answer
}
