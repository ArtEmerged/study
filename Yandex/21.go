package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(file)
	var n int
	fmt.Fscan(in, &n)
	f := make([]int, 36)
	i := 4
	//вычисляем кол-во последовательностей
	//сумма трех предыдущих элементов
	f[1], f[2], f[3] = 2, 4, 7
	for i <= n {
		f[i] = f[i-1] + f[i-2] + f[i-3]
		i++
	}
	fmt.Println(f[n])
}
