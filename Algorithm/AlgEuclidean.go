package main

func main() {

	print(GCD(1680, 640), "\n")
}
/*  The Euclidean algorithm is a way to find
the greatest common divisor(GCD) of two positive integers, 
a and b.*/
func GCD(x, y int) int {
	if x%y != 0 {
		y =GCD(y, x%y)
	}
	return y

}
