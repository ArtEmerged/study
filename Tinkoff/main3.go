// // #3
// package main

// import "fmt"

//	func main() {
//		n := 0
//		fmt.Scan(&n)
//		str1:=make([]int, n)
//		str1:=make([]int, n)
//	}

package main

import (
	"fmt"
)

// Функция canGetWinningSequence проверяет, можно ли получить выигрышную последовательность из a в b.
func canGetWinningSequence(n int, a, b []int) bool {
	firstNonZero := -1 // Индекс первого ненулевого элемента разницы
	lastNonZero := -1  // Индекс последнего ненулевого элемента разницы

	// Находим первый ненулевой элемент разницы
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			firstNonZero = i
			break
		}
	}

	// Находим последний ненулевой элемент разницы, начиная с конца
	for i := n - 1; i >= 0; i-- {
		if a[i] != b[i] {
			lastNonZero = i
			break
		}
	}

	// Если все элементы разницы равны нулю, то последовательность уже выигрышная
	if firstNonZero == -1 {
		return true
	}

	// Вычисляем разницу для проверки упорядочивания подотрезка [firstNonZero, lastNonZero]
	diff := a[firstNonZero] - b[firstNonZero]
	for i := firstNonZero; i <= lastNonZero; i++ {
		if a[i]-b[i] != diff {
			return false
		}
	}

	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	b := make([]int, n)

	// Считываем последовательности a и b
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	// Проверяем, можно ли получить выигрышную последовательность и выводим результат
	if canGetWinningSequence(n, a, b) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

