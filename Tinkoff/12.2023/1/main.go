package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		tinkoff := map[rune]int{
			'T': 1,
			'I': 1,
			'N': 1,
			'K': 1,
			'O': 1,
			'F': 2,
		}
		var word string
		fmt.Scan(&word)
		for _, letter := range word {
			tinkoff[letter]--
		}
		for _, count := range tinkoff {
			if count != 0 {
				res[i] = "No"
				break
			}
		}
		if res[i] == "" {
			res[i] = "Yes"
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
