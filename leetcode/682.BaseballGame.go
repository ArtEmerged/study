package leet

import "strconv"

func calPoints(op []string) int {
	arr := make([]int, 0, len(op))
	answer := 0
	for _, v := range op {
		switch v {
		case "C":
			if len(arr) > 0 {
				answer -= arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			}
		case "D":
			if len(arr) > 0 {
				arr = append(arr, 2*arr[len(arr)-1])
				answer += arr[len(arr)-1]
			}
		case "+":
			if len(arr) > 1 {
				arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])
				answer += arr[len(arr)-1]
			}
		default:
			i, _ := strconv.Atoi(v)
			arr = append(arr, i)
			answer += i
		}
	}
	return answer
}
