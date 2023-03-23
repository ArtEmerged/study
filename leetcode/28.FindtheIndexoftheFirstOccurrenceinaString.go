package leete

func strStr(s string, find string) int {
	for i := 0; i < len(s)-len(find)+1; i++ {
		for j := 0; j < len(find); j++ {
			if s[j+i] != find[j] {
				break
			} else if j == len(find)-1 {
				return i
			}
		}
	}
	return -1
}
