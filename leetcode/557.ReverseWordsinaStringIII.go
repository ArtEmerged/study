package leet

func reverseWords(s string) string {
    slice := strings.Split(s, "")
	copy := make([]string, len(slice))
	slice = append(slice, " ")
	m := 0
	for n := 0; n < len(slice); n++ {
		if slice[n] == " " {
			for i := n - 1; i >= m; i-- {
				copy = append(copy, slice[i])
			}
			m = n + 1
			copy = append(copy, " ")
		}
	}

	return strings.Join(copy[:len(copy)-1], "")
}