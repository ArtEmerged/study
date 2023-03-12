package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	A := convert()
	B := convert()
	C := convert()
	if C[0] < A[0] {
		C[0] += 24
	}
	a := (A[0]*60+A[1])*60 + A[2]
	b := (B[0]*60+B[1])*60 + B[2]
	c := (C[0]*60+C[1])*60 + C[2]

	answer := int(math.Ceil(float64(c-a)/2)) + b

	h := answer / 3600
	m := (answer % 3600) / 60
	s := answer % 60
	if answer/3600 > 23 {
		h = answer/3600 - 24
	}
	fmt.Printf("%02d:%02d:%02d\n", h, m, s)
}

func convert() []int {
	var ABC string
	result := make([]int, 3)
	fmt.Scanf("%s\n", &ABC)
	time := strings.Split(ABC, ":")
	for i, v := range time {
		for _, r := range v {
			result[i] = result[i]*10 + int(r-48)
		}
	}
	return result
}
