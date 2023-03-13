package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {

	var n, m int
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)
	fmt.Fscan(in, &n, &m)

	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
		for j := range grid[i] {
			fmt.Fscan(in, &grid[i][j])
		}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	fmt.Println(dp[n-1][m-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
