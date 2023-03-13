package main

import (
	"fmt"
)

func heapify(arr []int, n int, i int) {
	largest := i     // Инициализируем наибольший элемент как корень
	left := 2*i + 1  // Левый потомок корня
	right := 2*i + 2 // Правый потомок корня
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	heapSort(arr)
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}
