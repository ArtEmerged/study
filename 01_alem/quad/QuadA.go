package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if j == 1 && i == 1 {
					z01.PrintRune('o') // left high corner
				} else if j == x && i == 1 {
					z01.PrintRune('o') // right high corner
				} else if j == 1 && i == y {
					z01.PrintRune('o') // left low corner
				} else if j == x && i == y {
					z01.PrintRune('o') // right low corner
				} else if i == y || i == 1 {
					z01.PrintRune('-') // horizontally
				} else if j == x || j == 1 {
					z01.PrintRune('|') // vertical
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
