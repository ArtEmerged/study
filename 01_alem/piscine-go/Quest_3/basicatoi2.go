package piscine

func BasicAtoi2(s string) int {
	Int := 0
	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		}
		i = i - 48
		Int = Int*10 + int(i)
	}
	return Int
}
