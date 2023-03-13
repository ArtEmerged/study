package piscine

func ConcatParams(args []string) string {
	var result string
	for s, i := range args {
		result += i
		if s != len(args)-1 {
			result += "\n"
		}
	}
	return result
}
