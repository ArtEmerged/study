package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f1, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	in := bufio.NewReader(f1)
	pl1 := make([]int, 5)
	pl2 := make([]int, 5)
	for i := 0; i < 10; i++ {
		if i < 5 {
			fmt.Fscan(in, &pl1[i])
		} else {
			fmt.Fscan(in, &pl2[i-5])
		}
	}
	account := 0
	for len(pl1) != 0 && len(pl2) != 0 {
		if account == 1000000 {
			fmt.Println("botva")
			return
		}
		if pl1[0] == 0 && pl2[0] == 9 {
			pl1 = append(pl1, pl1[0], pl2[0])
			pl1 = pl1[1:]
			pl2 = pl2[1:]
		} else if pl2[0] == 0 && pl1[0] == 9 {
			pl2 = append(pl2, pl1[0], pl2[0])
			pl1 = pl1[1:]
			pl2 = pl2[1:]
		} else if (pl1[0] == 0 && pl2[0] == 9) || (pl1[0] > pl2[0]) {
			pl1 = append(pl1, pl1[0], pl2[0])
			pl1 = pl1[1:]
			pl2 = pl2[1:]
		} else if (pl2[0] == 0 && pl1[0] == 9) || (pl1[0] < pl2[0]) {
			pl2 = append(pl2, pl1[0], pl2[0])
			pl1 = pl1[1:]
			pl2 = pl2[1:]
		}
		account++
	}

	if len(pl2) == 0 {
		fmt.Printf("first %d\n", account)
	} else {
		fmt.Printf("second %d\n", account)
	}
}
