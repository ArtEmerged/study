package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "P0"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s) && j > 0; i++ {
		s := strings.ToLower(s)
		for i < j-1 && !(s[i] >= 'a' && s[i] <= 'z') {
			i++
		}
		for j > i+1 && !(s[j] >= 'a' && s[j] <= 'z') {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
