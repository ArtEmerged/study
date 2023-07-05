package main

import "fmt"

func main() {
	var b byte = 38
	fmt.Println(b)
	result := int(ReverseBits(b))
	// Print Byte
	slice := make([]int, 8)
	i := len(slice) - 1
	for result != 0 {
		slice[i] = result % 2
		result /= 2
		i--
	}
	fmt.Println(slice)
	//for _, w := range slice {
	//	z01.PrintRune(rune(w) + 48)
	//}
	//z01.PrintRune('\n')
}

// func ReverseBits(oct byte) byte {
// 	var res byte
// 	for i := 8; i > 0; i-- {
// 		res = (res << 1) | (oct & 1)
// 		oct >>= 1
// 	}
// 	return res
// }

func ReverseBits(oct byte) byte {
	var res byte
	for i := 8; i > 0; i-- {
		res = (res << 1) | (oct & 1)
		oct >>= 1
	}
	return res
}
