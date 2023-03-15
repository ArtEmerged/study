package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) == 1 {
		fmt.Print("File name missing\n")
		return
	} else if len(arg) > 2 {
		fmt.Print("Too many arguments\n")
		return
	}
	r, _ := os.Open(arg[1])
	file, _ := ioutil.ReadAll(r)
	fmt.Print(string(file))
}
