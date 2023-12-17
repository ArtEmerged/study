package main

import "fmt"

func main() {
	var countData int
	fmt.Scan(&countData)
	res := make([]string, countData)
	for i := 0; i < countData; i++ {
		var n, sum int
		fmt.Scan(&n)
		for i2 := 0; i2 < n; i2++ {
			var soc int
			fmt.Scan(&soc)
			sum += soc

		}
		if sum/2 >= n-1 {
			res[i] = "Yes"
		} else {
			res[i] = "No"
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
