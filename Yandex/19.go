package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)
	lenght := 0
	fmt.Fscanln(in, &lenght)
	arr := make([]int, 0, lenght)
	for z := 0; z < lenght; z++ {
		step, num := 0, 0
		fmt.Fscanln(in, &step, &num)
		switch step {
		case 0: // add element
			add(&arr, num)
		case 1: //delete element
			fmt.Println(arr[0])
			delete(&arr)
		}
	}
}

func add(arr *[]int, num int) {
	*arr = append(*arr, num)
	last := len(*arr) - 1
	dad := (last - 1) / 2
	for last > 0 && (*arr)[last] > (*arr)[dad] {
		(*arr)[last], (*arr)[dad] = (*arr)[dad], (*arr)[last]
		last = dad
		dad = (last - 1) / 2
	}
}

func delete(arr *[]int) {
	last := len(*arr) - 1
	i := 0
	(*arr)[i] = (*arr)[last]
	for 2*i+2 < len(*arr) {
		max := 2*i + 1
		if (*arr)[max] < (*arr)[2*i+2] {
			max = 2*i + 2
		}
		if (*arr)[i] < (*arr)[max] {
			(*arr)[i], (*arr)[max] = (*arr)[max], (*arr)[i]
			i = max
		} else {
			break
		}
	}
	*arr = (*arr)[:last]
}
