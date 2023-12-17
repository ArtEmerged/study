package main

import (
	"bufio"
	"fmt"
	"io"
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
			continue
		}
		go serve(conn)
	}
}

func serve(c net.Conn) {
	buf := bufio.NewScanner(c)
	for buf.Scan() {
		line := buf.Text()
		fmt.Println(line)
		if line == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}
	defer c.Close()

	fmt.Println("Code got here.")
	io.WriteString(c, "I see you connected\n")

}
