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
	var lenght int
	fmt.Fscanln(in, &lenght)
	arr := make([]int, 0, lenght)
	for z := 0; z < lenght; z++ {
		var num int
		fmt.Fscan(in, &num)
		arr = append(arr, num)
		last := len(arr) - 1
		dad := (last - 1) / 2
		for len(arr) > 0 && arr[last] < arr[dad] {
			arr[last], arr[dad] = arr[dad], arr[last]
			last = dad
			dad = (last - 1) / 2
		}
	}
	for pos := len(arr) - 1; pos >= 0; pos-- {
		fmt.Printf("%d ", arr[0])
		buf := arr[0]
		arr[0] = arr[pos]
		max(&arr, pos)
		arr[pos] = buf
	}
	fmt.Println(arr)

}

func max(arr *[]int, pos int) {
	i := 0
	for 2*i+2 < pos+1 {
		min := 2*i + 1
		if (*arr)[min] > (*arr)[2*i+2] {
			min = 2*i + 2
		}
		if (*arr)[i] > (*arr)[min] {
			(*arr)[min], (*arr)[i] = (*arr)[i], (*arr)[min]
			i = min
		} else {
			break
		}
	}

}
