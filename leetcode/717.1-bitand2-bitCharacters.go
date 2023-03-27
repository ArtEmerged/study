package leet

func isOneBitCharacter(bits []int) bool {
	l := len(bits)

	if l == 1 && bits[0] == 0 {
		return true
	} else if l == 2 && bits[0] == 1 {
		return false
	}

	for i := 1; i < l-1; i++ {
		if bits[i-1] == 1 {
			bits[i], bits[i-1] = -1, -1
			i++
		}
	}

	if bits[l-2] == -1 {
		return true
	}

	return bits[l-2] != 1
}
