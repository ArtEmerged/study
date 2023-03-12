package main

import (
	"fmt"
	"math"
)

func main() {
	var K int 
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := 0, 0
	fmt.Scanln(&K)
	for i := 0; i < K; i++ {
		var x, y int
		fmt.Scanln(&x, &y)
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
	}
	fmt.Printf("%d %d %d %d\n", minX, minY, maxX, maxY)
}
