package main

import "fmt"

func main() {
	var k int
	var word string
	fmt.Scan(&k, &word)
	lenghtMAX := 0
	for letter := 'a'; letter <= 'z'; letter++ {
		swap, Right := k, 0
		for Left, v := range word {
			if Right == len(word)-1 {
				break
			}
			for Right < len(word) && swap >= 0 {
				if word[Right] != byte(letter) {
					swap--
				}
				if swap >= 0 {
					Right++
				}
			}
			if v != letter {
				swap += 2
			}
			if Right-Left > lenghtMAX {
				lenghtMAX = Right - Left
			}
		}
	}
	fmt.Println(lenghtMAX)
}
