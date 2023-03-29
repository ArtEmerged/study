package leet

func subtractProductAndSum(n int) int {
	Product, Sum := 1, 0
	for n > 0 {
		Sum += n % 10
		Product *= (n % 10)
		n /= 10
	}
	return Product - Sum
}
