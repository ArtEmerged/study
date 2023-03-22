package leet

import "fmt"

func addToArrayForm(num []int, k int) []int {
	var arrK []int
	for k > 0 {
		arrK = append(arrK, k%10)
		k /= 10
	}
for len
	ost := 0
	for i := len(arrK) - 1; i >= 0; i-- {
		sum := num[i] + arrK[i]
		if ost == 1 && sum+ost > 9 {
			arrK[i] = (sum) - 9
		} else if ost == 1 {
			arrK[i] = (sum) + ost
			ost--
		} else if sum > 9 {
			arrK[i] = (sum) - 10
			ost++
		} else {
			arrK[i] = sum
		}
	}
	if ost > 0 {
		answer := []int{ost}
		answer = append(answer, arrK...)
		return answer
	}
	return arrK
}

func main() {
	num := []int{1, 2}
	k := 99
	a := addToArrayForm(num, k)
	fmt.Println(a)
	test := make([]int, 3)
	fmt.Println(test)
	test = append(test, 5)
	fmt.Println(test)

}
