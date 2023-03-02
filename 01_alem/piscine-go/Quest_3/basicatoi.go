package piscine

func BasicAtoi(s string) int {
	copy := []rune(s)
	Int := 0
	for _, i := range copy {
		Int = Int*10 + int(i-48)
	}

	return Int
}
