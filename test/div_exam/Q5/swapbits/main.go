package main

import "fmt"

func main() {
//SwapBits
	var b byte = 65
	result := int(SwapBits(b))
	// Print Byte
	slice := make([]int, 8)
	i := len(slice) - 1
	for result != 0 {
		slice[i] = result % 2
		result /= 2
		i--
	}
	fmt.Println(slice)
}

func SwapBits(octet byte) byte {
	return octet<<4 + octet>>4
}

// func ReverseBits(oct byte) byte {

// }
