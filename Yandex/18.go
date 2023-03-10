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
	in := bufio.NewReader(f)
	arr := make([]int, 0, 100)
	for {
		step := ""
		num := 0
		fmt.Fscanln(in, &step, &num)
		last := len(arr) - 1

		switch step {
		case "push_front":
			add := []int{num}
			arr = append(add, arr...)
			fmt.Println("ok")
		case "push_back":
			arr = append(arr, num)
			fmt.Println("ok")
		case "pop_front":
			fmt.Println(arr[0])
			arr = arr[1:]
		case "pop_back":
			fmt.Println(arr[last])
			arr = arr[:last]
		case "front":
			fmt.Println(arr[0])
		case "back":
			fmt.Println(arr[last])
		case "size":
			fmt.Println(len(arr))
		case "clear":
			arr = arr[:0]
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			return
		}
	}

}
