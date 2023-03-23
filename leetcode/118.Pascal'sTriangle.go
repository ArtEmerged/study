package leet

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		v := make([]int, i+1)
		v[0] = 1
		v[i] = 1
		for j := 1; j < len(v)-1; j++ {
			v[j] = ans[i][j] + ans[i][j-1]
		}
		ans = append(ans, v)
	}
	return ans[1:]
}
