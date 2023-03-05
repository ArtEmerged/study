package piscine

func ToUpper(s string) string {
	array := []rune(s)

	for i := 0; i < len(array); i++ {
		if array[i] > 96 && array[i] < 123 {
			array[i] -= 32
		}
	}
	return string(array)
}
