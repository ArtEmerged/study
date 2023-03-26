package leet

func isSubsequence(s string, t string) bool {
	v, g := 0, 0
	str := ""
	for v < len(s) && g < len(t) {
		if s[v] == t[g] {
			v++
			str += string(t[g])
		}
		g++
	}
	return str == s
}
