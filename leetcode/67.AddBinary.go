package leet

func addBinary(a string, b string) string {
	for len(a) != len(b) {
		if len(a) > len(b) {
			b = "0" + b
		} else {
			a = "0" + a
		}
	}
	save := false
	var answer string
	for i := len(a) - 1; i >= 0; i-- {
		if save {
			if a[i] == '1' && b[i] == '1' {
				answer = "1" + answer
				save = true
			} else if a[i] != b[i] {
				answer = "0" + answer
			} else {
				answer = "1" + answer
				save = false
			}
		} else {
			if a[i] == '1' && b[i] == '1' {
				answer = "0" + answer
				save = true
			} else if a[i] != b[i] {
				answer = "1" + answer
			} else {
				answer = "0" + answer
			}
		}
	}
	if save {
		answer = "1" + answer
	}
	return answer
}
