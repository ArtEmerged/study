package leet

func isPalindrome(x int) bool {
	revX, tmp := x, 0
	for revX > 0 {
		tmp *=10 + revX%10
		revX /= 10
	}
	return tmp == x
}
