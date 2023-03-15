package piscine

func Map(f func(int) bool, a []int) []bool {
	var TrFl []bool
	for _, n := range a {
		TrFl = append(TrFl, f(n))
	}
	return TrFl
}
