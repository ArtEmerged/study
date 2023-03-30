package leet

import (
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	answer, _ := strconv.Atoi(strings.Replace(strconv.Itoa(num), "6", "9", 1))
	return answer
}
