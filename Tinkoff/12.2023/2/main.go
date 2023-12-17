package main

import "fmt"

func main() {
	var countData int
	fmt.Scan(&countData)
	res := make([]string, countData)
	for i := 0; i < countData; i++ {
		var developers, n int
		fmt.Scan(&developers)
		for i2 := 0; i2 < developers; i2++ {
			var soc int
			fmt.Scan(&soc)
			if soc == 1 {
				n++
			}
		}
		if n > 2 {
			res[i] = "No"
		} else {
			res[i] = "Yes"
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}

}
