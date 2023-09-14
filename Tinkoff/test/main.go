package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// Создаем граф.
	graph := make([][]int, n)
	for i := 0; i < m; i++ {
		var v, u, w int
		fmt.Scanf("%d %d %d", &v, &u, &w) // <-- здесь исправили ошибку
		graph[v-1] = append(graph[v-1], u-1)
		graph[u-1] = append(graph[u-1], v-1)
	}

	// Находим компоненты связности графа.
	cmp := make([]int, n)
	for i := 0; i < n; i++ {
		cmp[i] = i
	}
	for i := 0; i < n; i++ {
		for _, j := range graph[i] {
			if cmp[i] != cmp[j] {
				union(cmp, i, j)
			}
		}
	}

	// Находим минимальное значение x, при котором количество компонент связности не изменится.
	x := 0
	for _, w := range graph {
		for _, w := range w {
			w++
			if w <= x {
				x++
			}
		}
	}

	// Выводим ответ.
	fmt.Println(x)
}

func union(cmp []int, i, j int) {
	cmp[find(cmp, i)] = find(cmp, j)
}

func find(cmp []int, i int) int {
	if cmp[i] != i {
		cmp[i] = find(cmp, cmp[i])
	}
	return cmp[i]
}
