package piscine

func Swap(a, b *int) {
	*a, *b = *b, *a
}
