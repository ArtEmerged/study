package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		buf := bufio.NewScanner(conn)
		for buf.Scan() {
			line := buf.Text()
			fmt.Println(line)
		}
		defer conn.Close()

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected\n")
	}
}
