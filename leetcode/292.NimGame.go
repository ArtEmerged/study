package leet

func canWinNim(n int) bool {
	if n < 4 {
		return true
	}
	if n%4 == 0 {
		return false
	}
	return true
}
