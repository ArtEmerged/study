package piscine

func NRune(s string, n int) rune {
	res := []rune(s)

	if n > len(res) || n < 1 {
		return 0
	}
	return res[n-1]
}
