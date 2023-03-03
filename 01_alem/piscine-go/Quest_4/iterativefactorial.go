package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	} else if nb < 0 || nb > 20 {
		return 0
	}
	fact := nb
	for i := nb - 1; i > 0; i-- {
		fact *= i
	}
	return fact
}
