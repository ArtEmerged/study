package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		pow := 2
		n,_ := strconv.Atoi(os.Args[1])
		for i := 0; i < 12; i++ {
			if pow == n {
				for _, vol := range "true" {
					z01.PrintRune(vol)
				}
				z01.PrintRune(10)
				return
			}
			pow *= 2
		}
		for _, vol := range "false" {
			z01.PrintRune(vol)
		}
		z01.PrintRune(10)
		return
	}
}


// func main() {
// 	if len(os.Args) == 2 {
// 		n, _ := strconv.Atoi(os.Args[1])
// 		if n != 0 && n&(n-1) == 0 {
// 			for _, vol := range "true" {
// 				z01.PrintRune(vol)
// 			}
// 			z01.PrintRune(10)
// 			return
// 		}
// 		for _, vol := range "false" {
// 			z01.PrintRune(vol)
// 		}
// 		z01.PrintRune(10)
// 		return
// 	}
// }
