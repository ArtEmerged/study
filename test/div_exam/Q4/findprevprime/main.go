package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(33))
	fmt.Println(FindPrevPrime(100))
}

// func FindPrevPrime(nb int) int {
// 	if nb <= 1 {
// 		return 0
// 	}
// 	for i := nb; i > 0; i-- {
// 		check := true
// 		for val := 2; val < i; val++ {
// 			if i%val == 0 {
// 				check = false
// 				break
// 			}
// 		}
// 		if check {
// 			return i
// 		}
// 	}
// 	return 1
// }

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}
	for nb > 0 {
		if IsPrime(nb) {
			return nb
		}
		nb--
	}
	return 1
}

func IsPrime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
