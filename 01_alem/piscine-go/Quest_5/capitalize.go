package piscine

func Capitalize(s string) string {
	R := []rune(s)
	Lowercase(R)
	Capitalizes(R)
	return string(R)
}

func Lowercase(R []rune) {
	for i := 0; i < len(R); i++ {
		if R[i] >= 'A' && R[i] <= 'Z' {
			R[i] = R[i] + 32
		}
	}
}

func Capitalizes(R []rune) {
	for i := 0; i < len(R)-1; i++ {
		if ((R[i] < 'a' || R[i] > 'z') && (R[i] < '0' || R[i] > '9')) && (R[i+1] >= 'a' && R[i+1] <= 'z') {
			R[i+1] = R[i+1] - 32
			i++
		}
	}
	if R[0] >= 'a' && R[0] <= 'z' {
		R[0] = R[0] - 32
	}
}
