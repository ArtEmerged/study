package piscine

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) < 97 || rune(s[i]) > 122 {
			return false
		}
	}
	return true
}
