package piscine

func TrimAtoi(s string) int {
	var res int
	minus := 1
	for _, n := range s {
		if n == '-' && res == 0 {
			minus = -1
		}
		if n >= '0' && n <= '9' {
			res = res*10 + int(n-48)
		}
	}
	return res * minus
}
