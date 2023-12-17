package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	defer l.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	Request(c)
	fmt.Println("### METHOD:", m, "###")
	fmt.Println("### URL:", u, "###")
	ResponseWriter(c)
}

var m, u, body string

func Request(r net.Conn) {
	scanner := bufio.NewScanner(r)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if i == 0 {
			m = strings.Fields(line)[0]
			u = strings.Fields(line)[1]
		}
		if line == "" {
			break
		}
	}

}

func ResponseWriter(w net.Conn) {
	body += `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>index</title>
	</head>
	<body>
		<h1>HOLY COW THIS IS LOW LEVEL</h1>
	</body>
	</html>`
	fmt.Fprint(w, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(w, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(w, "Content-Type: text/html\r\n")
	fmt.Fprint(w, "\n\r")
	fmt.Fprint(w, body)
}
