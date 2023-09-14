package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	count := 0
	for i := 1.; i < 19; i++ {
		for j := 1; j <= 9; j++ {
			polindrom := j * (int(math.Pow(10, i)-1) / (10 - 1))
			if polindrom > b {
				break
			}
			if polindrom >= a {
				count++
			}
		}
	}
	fmt.Println(count)
}
