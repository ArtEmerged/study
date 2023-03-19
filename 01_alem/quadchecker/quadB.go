package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	x, _ := strconv.Atoi(arg1)
	y, _ := strconv.Atoi(arg2)
	QuadB(x, y)
}

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if j == 1 && i == 1 {
					fmt.Print("/") // left high corner
				} else if j == x && i == 1 {
					fmt.Print("\\") // right high corner
				} else if j == 1 && i == y {
					fmt.Print("\\") // left low corner
				} else if j == x && i == y {
					fmt.Print("/") // right low corner
				} else if i == y || i == 1 {
					fmt.Print("*") // horizontally
				} else if j == x || j == 1 {
					fmt.Print("*") // vertical
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}
