package piscine

func MakeRange(min, max int) []int {
	if min > max || min == max {
		return nil
	}
	slice := make([]int, max-min)
	for i, s := 0, min; i < len(slice); i++ {
		slice[i] = s
		s++
	}
	return slice
}
