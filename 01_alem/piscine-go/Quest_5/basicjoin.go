package piscine

func BasicJoin(elems []string) string {
	var answer string
	for _, n := range elems {
		answer += n
	}
	return answer
}
