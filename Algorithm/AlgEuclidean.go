package main

func main() {

	print(GCD(1680, 640), "\n")
}

func GCD(x, y int) int {
	if x%y != 0 {
		y =GCD(y, x%y)
	}
	return y

}
