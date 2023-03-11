package main

import "fmt"

func main() {
	var N, K, y, x, place int
	fmt.Scanf("%d\n %d\n %d\n %d", &N, &K, &y, &x)
	if x == 1 || K%2 == 0 {
		if x == 1 {
			place = y*2 - 1
		} else {
			place = y * 2
		}
		if N-place >= K {
			print(place + K)
		} else if place-K > 0 {
			print(place - K)
		} else {
			fmt.Println(-1)
		}
	} else {
		place = y * 2
		if place-K > 0 {
			print(place - K)
		} else if N-place >= K {
			print(place + K)
		} else {
			fmt.Println(-1)
		}
	}
}

func print(place2 int) {
	if place2%2 == 0 {
		fmt.Printf("%d %d\n", place2/2, 2)
	} else {
		fmt.Printf("%d %d\n", (place2+1)/2, 1)
	}
}
