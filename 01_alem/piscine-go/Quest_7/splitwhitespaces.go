package piscine

func SplitWhiteSpaces(s string) []string {
	str := ""
	s += string(' ')
	var array []string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			if len(str) > 0 {
				array = append(array, str)
			}
			str = ""
		} else {
			str += string(s[i])
		}
	}
	return array
}
