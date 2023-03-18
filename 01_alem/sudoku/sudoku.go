package main

import (
	"os"

	"github.com/01-edu/z01"
)

var sign = false

var sudoku [9][9]int

func main() {
	arr := os.Args[1:]
	err := "Error"
	if checkErrors(arr) {
		for _, i := range err {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	} else {
		optimizeSudoku(arr)
		DFS(0)
		Output(sudoku)
	}
}

func Check(n, key int) bool {
	for i := 0; i < 9; i++ { ///cheak row
		r := n / 9

		if sudoku[r][i] == key {
			return false
		}
	}
	for i := 0; i < 9; i++ { ///cheak column
		c := n % 9
		if sudoku[i][c] == key {
			return false
		}
	}
	x := n / 9 / 3 * 3
	y := n % 9 / 3 * 3
	for i := x; i < x+3; i++ { ///cheak quad
		for j := y; j < y+3; j++ {
			if sudoku[i][j] == key {
				return false
			}
		}
	}
	return true
}

func DFS(n int) {
	if n > 80 {
		sign = true
		return
	}
	if sudoku[n/9][n%9] != 0 { // search zero in Arr
		DFS(n + 1)
	} else {
		for key := 1; key <= 9; key++ {
			if Check(n, key) {
				sudoku[n/9][n%9] = key
				DFS(n + 1)
				if sign {
					return
				}
				sudoku[n/9][n%9] = 0
			}
		}
	}
}

func checkErrors(arr []string) bool {
	if len(arr) != 9 {
		return true
	}
	for i, str := range arr {
		if len(str) != 9 {
			return true
		}
		for _, ch := range arr[i] {
			if !(ch >= '1' && ch <= '9' || ch == '.') {
				return true
			}
		}
	}
	return false
}

func optimizeSudoku(arr []string) {
	for i, str := range arr {
		for j, ch := range str {
			if ch != '.' {
				sudoku[i][j] = int(ch - 48)
			}
		}
	}
}

func Output(arr [9][9]int) {
	err := "Error"

	if Errors(arr) {
		for _, i := range err {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				z01.PrintRune(rune(arr[i][j]) + 48)
				if j != len(arr)-1 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func Errors(arr [9][9]int) bool {
	for d := 0; d < len(arr); d++ {
		for f := 0; f < len(arr); f++ {
			if arr[d][f] == 0 {
				return true
			}
		}
	}
	return false
}
