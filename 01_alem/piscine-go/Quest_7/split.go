package piscine

func Split(s, sep string) []string {
	var answ []string
	s += sep
	var cat string
	lSep := len(sep)
	for i := 0; i < len(s); i++ {
		if s[i:lSep+i] == sep {
			if cat > "" {
				answ = append(answ, cat)
			}
			cat = ""
			i += lSep - 1
			continue

		}
		cat += string(s[i])
	}
	return answ
}
