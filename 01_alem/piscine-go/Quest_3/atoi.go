package piscine

func Atoi(s string) int {
	New := s
	Int := 0
	Minus := 1

	if len(s) > 1 {
		if s[0] == '+' {
			New = s[1:]
		} else if s[0] == '-' {
			Minus = -1
			New = s[1:]
		}
	}
	for _, i := range New {
		if i < '0' || i > '9' {
			return 0
		}
		i = i - 48
		Int = Int*10 + int(i)
	}
	return Int * Minus
}
