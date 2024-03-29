package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		li := scan.Text()
		fmt.Println(li)
		fmt.Fprintf(conn, "I heard you say: %s\n", li)
	}
	defer conn.Close()
	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}
