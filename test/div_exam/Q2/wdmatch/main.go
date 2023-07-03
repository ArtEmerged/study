package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) == 3 {
// 		a1 := os.Args[1]
// 		a2 := []rune(os.Args[2])
// 		pos := 0

// 		for _, n := range a1 {
// 			if b, g := check(n, a2, pos); b {
// 				pos = g + 1
// 				continue
// 			} else {
// 				z01.PrintRune('\n')
// 				return
// 			}
// 		}
// 		for _, a := range a1 {
// 			z01.PrintRune(a)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func check(r rune, arr []rune, pos int) (bool, int) {
// 	for i := pos; i < len(arr); i++ {
// 		if r == arr[i] {
// 			return true, i
// 		}
// 	}
// 	return false, 0
// }


func main() {
	if len(os.Args) == 3 {
		arg := os.Args[1:]
		for i, j := 0, 0; i < len(arg[0]) && j < len(arg[1]); {
			if arg[0][i] == arg[1][j] {
				i++
				if i == len(arg[0]) {
					for _, w := range arg[0] {
						z01.PrintRune(w)
					}
					z01.PrintRune('\n')
				}
			}
			j++
			if j == len(arg[1]) {
				return
			}
		}
	}
}
