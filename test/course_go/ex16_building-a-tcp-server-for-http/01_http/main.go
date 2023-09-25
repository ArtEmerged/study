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

var li string

func request(r net.Conn) {
	scan := bufio.NewScanner(r)
	i := 0
	for scan.Scan() {
		if i == 0 {
			li = scan.Text()
			li = strings.Fields(li)[0]
		} else {
			break
		}
		i++
	}
	fmt.Println(li)
}

func response(w net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>` + li + `</title>
</head>
<body>
    <h1>` + li + `</h1>
</body>
</html>`
	fmt.Fprint(w, body)
}
