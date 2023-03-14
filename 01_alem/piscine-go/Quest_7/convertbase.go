package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	Lto := len(baseTo)
	Lfrom := len(baseFrom)
	var answer2 string
	if Lfrom < Lto {
		answer1 := Up(nbr, baseFrom)
		answer2 = Low(answer1, baseTo)
	} else {
		answer1 := atoi(nbr)
		answer2 = Low(answer1, baseTo)
	}
	return answer2
}

func Up(nbr, baseFrom string) int {
	Lnbr := len(nbr)
	var answer1 int
	for i, n := 0, Lnbr-1; i < Lnbr; i++ {
		pow := Pow1(len(baseFrom), n)
		answer1 += int(nbr[i]-'0') * pow
		n--
	}
	return answer1
}

func Low(nbr int, base string) (answer string) {
	b := len(base)
	var bezrev []rune
	for nbr > 0 {
		bezrev = append(bezrev, rune(base[nbr%b]))
		nbr /= b
	}
	for i := len(bezrev) - 1; i >= 0; i-- {
		answer += string(bezrev[i])
	}
	return answer
}

func Pow1(prime, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	}
	return prime * Pow1(prime, power-1)
}

func atoi(s string) (answ int) {
	for _, n := range s {
		answ = answ*10 + int(n-'0')
	}
	return answ
}
