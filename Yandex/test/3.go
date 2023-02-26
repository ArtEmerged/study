package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var quantitySticks, quantityCollector int
	var N, K string
	file, _ := os.Open("input.txt")
	defer file.Close()
	read := bufio.NewScanner(file)
	for read.Scan() {
		if quantitySticks == 0 {
			quantitySticks, _ = strconv.Atoi(read.Text())
		} else if N == "" {
			N = read.Text()
		} else if quantityCollector == 0 {
			quantityCollector, _ = strconv.Atoi(read.Text())
		} else {
			K = read.Text()
		}
	}
	arrN := strings.Split(N, " ")
	arrK := strings.Split(K, " ")
	StickersNumber := make([]int, 0, quantitySticks)
	StickersCollector := make([]int, 0, quantityCollector)
	key := map[string]int{}
	for _, s := range arrN {
		key[s]++
		if key[s] == 1 {
			num, err := strconv.Atoi(s)
			if err == nil {
				StickersNumber = append(StickersNumber, num)
			}
		}
	}
	for _, s := range arrK {
		num, err := strconv.Atoi(s)
		if err == nil {
			StickersCollector = append(StickersCollector, num)
		}
	}
	sortInt(StickersNumber)
	for _, number := range StickersCollector {
		fmt.Println(Search(StickersNumber, number))
	}
}

func sortInt(arr []int) []int {
	if len(arr) > 1 {
		Left := 0
		Right := len(arr) - 1
		middle := arr[(Right+Left)/2]
		for Left < Right {
			for arr[Left] < middle {
				Left++
			}
			for arr[Right] > middle {
				Right--
			}
			if Left <= Right {
				arr[Left], arr[Right] = arr[Right], arr[Left]
				Left++
				Right--
			}
		}
		sortInt(arr[:Right+1])
		sortInt(arr[Left:])
	}
	return arr
}

func Search(arr []int, find int) int {
	if find != 0 {
		Left := 0
		Right := len(arr) - 1
		answer := 0
		for Left <= Right {
			middle := (Right + Left) / 2
			if find == arr[middle] {
				return middle
			}
			if find > arr[middle] {
				Left = middle + 1
			} else {
				Right = middle - 1
			}
			answer = middle
		}
		if find > arr[answer] {
			return answer + 1
		}
		return answer
	}
	return 0
}
