package piscine

func Index(s string, toFind string) int {
	for st := 0; st < len(s)-len(toFind)+1; st++ {
		if toFind == s[st:len(toFind)+st] {
			return st
		}
	}
	return -1
}

func findMatch(s, toFind string) int {
	for i := 0; i < len(s)-len(toFind)+1; i++ {
		for j := 0; j < len(toFind); j++ {
			if s[j+i] != toFind[j] {
				break
			} else if j == len(toFind)-1 {
				return i
			}
		}
	}
	return -1
}
