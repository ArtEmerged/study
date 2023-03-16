package piscine

func CountIf(f func(string) bool, tab []string) int {
	Answer := 0
	for _, n := range tab {
		if f(n) {
			Answer++
		}
	}
	return Answer
}
