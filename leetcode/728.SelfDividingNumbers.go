package leet

func selfDividingNumbers(left int, right int) []int {
	var res []int

	for i := left; i <= right; i++ {
		isEligible := false
		ct := i
		for ct != 0 {
			isEligible = true
			ctx := ct % 10
			if ctx == 0 || i%ctx != 0 {
				isEligible = false
				break
			}
			ct = ct / 10
		}
		if isEligible {
			res = append(res, i)
		}
	}
	return res
}
