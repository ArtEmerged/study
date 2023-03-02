package piscine

func StrRev(s string) string {
	var res string
	for _, i := range s {
		res = string(i) + res
	}
	return res
}
