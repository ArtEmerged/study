package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
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
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}
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
	fmt.Println("**CODE GOT HERE**")
}
