package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	defer conn.Close()

	request(conn)
	response(conn)
}

var res string

func request(r net.Conn) {
	scan := bufio.NewScanner(r)
	i := 0
	for scan.Scan() {
		li := scan.Text()
		fmt.Println(li)
		if i == 0 {
			m := strings.Fields(li)[1]
			res = fmt.Sprintf("*** URL: %s ***\n", m)
			fmt.Println(res)

		}
		if li == "" {
			break
		}
		i++
	}
}

func response(w net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>` + res + `</title>
	</head>
	<body>
    <h1>` + res + `</h1>
	</body>
	</html>`
	fmt.Fprint(w, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(w, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(w, "Content-Type: text/html\r\n")
	fmt.Fprint(w, "\r\n")
	fmt.Fprint(w, body)
}
