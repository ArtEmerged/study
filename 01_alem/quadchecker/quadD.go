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
	QuadD(x, y)
}

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if j == 1 && i == 1 {
					fmt.Print("A") // left high corner
				} else if j == x && i == 1 {
					fmt.Print("C") // right high corner
				} else if j == 1 && i == y {
					fmt.Print("A") // left low corner
				} else if j == x && i == y {
					fmt.Print("C") // right low corner
				} else if i == y || i == 1 {
					fmt.Print("B") // horizontally
				} else if j == x || j == 1 {
					fmt.Print("B") // vertical
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}
