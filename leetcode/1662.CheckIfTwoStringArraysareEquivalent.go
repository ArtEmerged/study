package leet

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var sumWord1, sumWord2 string
	for _, v := range word1 {
		sumWord1 += v
	}
	for _, v := range word2 {
		sumWord2 += v
	}
	return sumWord1 == sumWord2
}
