package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	check := false
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			check = true
			break
		}
	}
	if check {
		for i := nb + 1; check; i++ {
			for n := 2; n < nb; n++ {
				if i%n == 0 {
					break
				} else if n == nb-1 {
					check = false
					nb = i
				}
			}
		}
	}
	return nb
}
