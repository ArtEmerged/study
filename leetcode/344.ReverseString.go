package leet

func reverseString(s []byte) {
	i := 1
	for n := 0; n < len(s)/2; n++ {
		s[n], s[len(s)-i] = s[len(s)-i], s[n]
		i++
	}
}
