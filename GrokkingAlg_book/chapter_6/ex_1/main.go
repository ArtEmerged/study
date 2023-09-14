package main

import (
	"fmt"
)

var Friends = map[string][]string{
	"you":    {"alice", "claire", "bob"},
	"bob":    {"anuj", "peggy"},
	"alice":  {"peggy"},
	"claire": {"thom", "jonnH"},
	"anuj":   {},
	"peggy":  {},
	"thom":   {},
	"jonny":  {},
}

func main() {
	Queue("you")
}

var Check []string

type queue []string

func (q *queue) Push(i any) {
	switch i := i.(type) {
	case string:
		*q = append(*q, i)
	case []string:
		*q = append(*q, i...)
	}
}
func (q queue) Len() int {
	return len(q)
}
func (q *queue) Pop() string {
	old := *q
	if len(old) > 0 {
		x := old[0]
		(*q) = (*q)[1:]
		return x
	}
	return ""
}

// queue implementation in golang
func Queue(start string) {
	// queue := []string{}
	// queue = append(queue, Friends[start]...)
	q := &queue{start}
	for q.Len() != 0 {
		//Вытаскиваем первого человека из очереди
		person := q.Pop()
		//Двигаем очередь
		// queue = queue[1:]
		//Проверка был ли этот человек уже в очереди
		//В данном задаче мы добавляем наших друзей в очередь,
		//но у нас могут быть общие друзья и чтобы не добавлять их повторно в очередь мы делаем проверку,
		//Если это проверку не делать то мы можем упереться в бесконечный цикл когда У меня есть друг Дима а у Димы есть в списке друзей Я
		// Петя <---> Дима
		if CheckPerson(person) {
			//Проверяем является ли данный человек продавцом манго
			if Search(person) {
				//Если он продает манго, то возвращаем в stdout текст и прекращаем функцию
				fmt.Printf("%s is a mango seller!\n", person)
				return
			} else {
				//Добавляем человека в список кто уже стоял в очереди
				Check = append(Check, person)
				//Добавляем в очередь всех друзей person
				q.Push(Friends[person])
				//Добавляем к длине кол-во друзей
			}
		}
	}
	fmt.Printf("Empty :(\n")
}

func Search(s string) bool {
	return s[len(s)-1] == 'H'
}

func CheckPerson(s string) bool {
	for _, v := range Check {
		if v == s {
			return false
		}
	}
	return true
}
