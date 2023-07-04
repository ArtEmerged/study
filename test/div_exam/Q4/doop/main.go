package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		arg := os.Args[1:]
		a := Atoi(arg[0])
		b := Atoi(arg[2])
		if len(arg[0]) >= 19 || len(arg[2]) >= 19 {
			return
		}

		switch arg[1] {
		case "+":
			fmt.Printf("%d\n", a+b)
		case "-":
			fmt.Printf("%d\n", a-b)
		case "*":
			fmt.Printf("%d\n", a*b)
		case "/":
			if b == 0 {
				fmt.Println("No division by 0")
				return
			}
			fmt.Printf("%d\n", a/b)
		case "%":
			if b == 0 {
				fmt.Println("No division by 0")
				return
			}
			fmt.Printf("%d\n", a%b)
		}
	}
}

func Atoi(s string) int {
	plus_minus := 1

	if s[0] == '-' {
		plus_minus = -1
		s = s[1:]
	}
	res := 0
	for _, n := range s {
		if n < '0' || n > '9' {
			os.Exit(0)
		}
		res = res*10 + int(n-48)
	}
	return res * plus_minus
}
