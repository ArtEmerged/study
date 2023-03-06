package piscine

func ToLower(s string) string {
	array := []rune(s)

	for i := 0; i < len(array); i++ {
		if array[i] > 64 && array[i] < 91 {
			array[i] += 32
		}
	}
	return string(array)
}
