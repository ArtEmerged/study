package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	s, _ := os.ReadFile("input.txt")
	letMax := 0
	karta := map[string]int{}
	for _, v := range s {
		st := string(v)
		if v < 33 || v > 126 {
			continue
		}
		if _, kay := karta[st]; kay {
			karta[st] += 1
			continue
		}
		karta[st] = 1
	}

	sortKarta := make([]string, 0, len(karta))
	for i, n := range karta {
		sortKarta = append(sortKarta, i)
		if n > letMax {
			letMax = n
		}
	}
	sort.Strings(sortKarta)

	for line := letMax; line > 0; line-- {
		for _, v := range sortKarta {
			if karta[v] < line {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
	for _, i := range sortKarta {
		fmt.Print(i)
	}
	fmt.Print("\n")

}
