package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	defer li.Close()
	if err != nil {
		log.Fatalln(err)
	}
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
	// scan.Split(bufio.ScanLines)
	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
	fmt.Println("Code got here.")
}
