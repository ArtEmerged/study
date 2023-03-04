package piscine

func AlphaCount(s string) int {
	sum := 0
	alph := []rune(s)
	for i := 0; i < len(alph); i++ {
		if (alph[i] > 64 && alph[i] < 91) || (alph[i] > 96 && alph[i] < 123) {
			sum++
		}
	}
	return sum
}
