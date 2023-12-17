package main

import (
	"fmt"
)

const MAXN = 3e5 + 13

var (
	n, k int
	a    [MAXN]int64
	ca   [MAXN]string
	c    [MAXN]string
	sum  [MAXN]int64
	ans  int64
)

var (
	g    [MAXN][]int64
	need [MAXN][35]bool
)

func find_id(name string) int {
	for i := 1; i <= k; i++ {
		if name == c[i] {
			return i
		}
	}
	return -1
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func dfs(v int64) {
	sum[v] = a[v]
	currID := find_id(ca[v])
	need[v][currID] = true
	for _, to := range g[v] {
		dfs(to)
		sum[v] += sum[to]
		for i := 1; i <= k; i++ {
			need[v][i] = (need[v][i] || need[to][i])
		}
	}
	cnt := 0
	for i := 1; i <= k; i++ {
		if need[v][i] {
			cnt++
		}
	}
	if cnt == k {
		ans = min(ans, sum[v])
	}
}

func main() {
	fmt.Scan(&n, &k)
	for i := 1; i <= k; i++ {
		fmt.Scan(&c[i])
	}
	var root int64 = -1
	for i := 1; i <= n; i++ {
		var pr, cost int64
		var name string
		fmt.Scanf("%d %d %s", &pr, &cost, &name)
		if pr == 0 {
			root = int64(i)
		} else {
			g[pr] = append(g[pr], int64(i))
		}
		a[i] = cost
		ca[i] = name
	}
	ans = 3e9 + 10
	dfs(root)
	fmt.Println(ans)
}
