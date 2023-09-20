package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		li := scan.Text()
		fmt.Fprintln(conn, li)
	}
}
