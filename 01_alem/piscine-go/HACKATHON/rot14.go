package piscine

func Rot14(s string) string {
	R := []rune(s)
	for i, v := range R {
		if v >= 'a' && v <= 'm' || v >= 'A' && v <= 'M' {
			R[i] = v + rune(14)
		} else if v >= 'n' && v <= 'z' || v >= 'N' && v <= 'Z' {
			R[i] = v - rune(12)
		}
	}

	return string(R)
}