package main

import (
	"bufio"
	"fmt"
	"os"
)

func determine(table []string) (x, y int) {
	x = 0
	y = 0
	for i := range table {
		x = i
		for j := range table[0] {
			y = j
		}
	}
	return y + 1, x + 1
}

func main() {
	check := bufio.NewScanner(os.Stdin)
	stroke := ""
	var table []string
	for check.Scan() {
		stroke = stroke + check.Text()
		table = append(table, stroke)
		stroke = ""
	}
	var x, y int
	if len(table) > 0 {
		x, y = determine(table)
	} else {
		fmt.Println("Not a quad function")
		return
	}

	flag := true

	for _, elemtab := range table { // check len elements for similar
		counter := elemtab
		for _, lentab := range table {
			if len(counter) != len(lentab) {
				flag = false
				break
			}
		}
	}

	for i := 0; i < len(table); i++ { // check top
		for j := 1; j < len(table[i])-1; j++ {
			if table[i][j] != '-' && table[i][j] != ' ' && table[i][j] != '*' && table[i][j] != 'B' {
				flag = false
			}
		}
	}

	for i := 1; i < len(table)-1; i++ { // check first end last element in middle
		if table[i][0] != '|' && table[i][len(table[i])-1] != '|' && table[i][0] != '*' && table[i][len(table[i])-1] != '*' && table[i][0] != 'B' && table[i][len(table[i])-1] != 'B' {
			flag = false
		}
	}
	// x = number of simbol in row [][x]
	// y = number of row [y][]

	if flag == true {
		if x == 1 && y == 1 {
			OneToOne(table, x, y)
		} else if y == 1 && x > 1 {
			Horizont(table, x, y)
		} else if y > 1 && x == 1 {
			Vertical(table, x, y)
		} else if x > 1 && y > 1 {
			Square(table, x, y)
		}
	} else {
		fmt.Println("Not a quad function")
	}
}

func OneToOne(table []string, x, y int) {
	switch table[0][0] {
	case 'o':
		fmt.Println("[quadA] [1] [1]")
	case '/':
		fmt.Println("[quadB] [1] [1]")
	case 'A':
		fmt.Println("[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]")
	default:
		fmt.Println("Not a quad function")
	}
}

func Horizont(table []string, x, y int) {
	if table[0][0] == 'o' && table[0][x-1] == 'o' {
		fmt.Printf("[quadA] [%v] [%v]\n", x, y)
	} else if table[0][0] == '/' && table[0][x-1] == '\\' {
		fmt.Printf("[quadB] [%v] [%v]\n", x, y)
	} else if table[0][0] == 'A' && table[0][x-1] == 'A' {
		fmt.Printf("[quadC] [%v] [%v]\n", x, y)
	} else if table[0][0] == 'A' && table[0][x-1] == 'C' {
		fmt.Printf("[quadD] [%v] [%v]", x, y)
		fmt.Printf(" || [quadE] [%v] [%v]\n", x, y)
	} else {
		fmt.Println("Not a quad function")
	}
}

func Vertical(table []string, x, y int) {
	if table[0][0] == 'o' && table[y-1][0] == 'o' {
		fmt.Printf("[quadA] [%v] [%v]\n", x, y)
	} else if table[0][0] == '/' && table[y-1][0] == '\\' {
		fmt.Printf("[quadB] [%v] [%v]\n", x, y)
	} else if table[0][0] == 'A' && table[y-1][0] == 'A' {
		fmt.Printf("[quadD] [%v] [%v]\n", x, y)
	} else if table[0][0] == 'A' && table[y-1][0] == 'C' {
		fmt.Printf("[quadC] [%v] [%v]", x, y)
		fmt.Printf(" || [quadE] [%v] [%v]\n", x, y)
	} else {
		fmt.Println("Not a quad function")
	}
}

func Square(table []string, x, y int) {
	if table[0][0] == 'o' && table[0][x-1] == 'o' && table[y-1][0] == 'o' && table[y-1][x-1] == 'o' {
		fmt.Printf("[quadA] [%v] [%v]\n", x, y)
	} else if table[0][0] == '/' && table[0][x-1] == '\\' && table[y-1][0] == '\\' && table[y-1][x-1] == '/' {
		fmt.Printf("[quadB] [%v] [%v]\n", x, y)
	} else if table[0][0] == 'A' && table[0][x-1] == 'A' && table[y-1][0] == 'C' && table[y-1][x-1] == 'C' {
		fmt.Printf("[quadC] [%v] [%v]\n", x, y)
	} else if table[0][0] == 'A' && table[0][x-1] == 'C' && table[y-1][0] == 'A' && table[y-1][x-1] == 'C' {
		fmt.Printf("[quadD] [%v] [%v]\n", x, y)
	} else if table[0][0] == 'A' && table[0][x-1] == 'C' && table[y-1][0] == 'C' && table[y-1][x-1] == 'A' {
		fmt.Printf("[quadE] [%v] [%v]\n", x, y)
	} else {
		fmt.Println("Not a quad function")
	}
}
