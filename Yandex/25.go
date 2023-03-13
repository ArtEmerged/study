// // package main

// // import (
// // 	"bufio"
// // 	"fmt"
// // 	"log"
// // 	"os"
// // 	"sort"
// // )

// func main() {
// 	var n int
// 	f, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	in := bufio.NewReader(f)
// 	fmt.Fscan(in, &n)

// 	a := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(in, &a[i])
// 	}
// 	sort.Ints(a)

// 	ans := 0
// 	for i := 0; i < n-1; i++ {
// 		ans += a[i+1] - a[i]
// 	}
// 	fmt.Println(ans)
// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	var arr [101]int
	var dp [101]int
	var n int
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	sort.Ints(arr[1 : n+1])
	dp[2] = arr[2] - arr[1]
	dp[3] = arr[3] - arr[1]
	for i := 4; i <= n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + arr[i] - arr[i-1]
	}
	fmt.Println(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
