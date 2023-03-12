package main

import (
	"fmt"
	"sort"
)

func main() {
	var word string
	fmt.Scan(&word)
	letter := make(map[string]int)
	for i, v := range word {
		letter[string(v)] += (len(word) - i) * (i + 1)
	}
	sortLetter := make([]string, 0, len(letter))
	for i := range letter {
		sortLetter = append(sortLetter, string(i))
	} 
	sort.Strings(sortLetter)
	for _, v := range sortLetter {
		fmt.Printf("%s: %d\n", v, letter[v])
	}
}
