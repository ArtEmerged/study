package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == 1 {
			return false
		}
	}
	return true
}

func f(a, b int) int {
	if a > b {
		return 1
	} else {
		return 0
	}

}
