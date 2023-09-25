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

var (
	m string
	u string
)

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
	// response(conn)
}

func request(r net.Conn) {
	i := 0
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		li := scan.Text()
		fmt.Println(li)
		if i == 0 {
			mux(r, li)
		}
		if li == "" {
			break
		}
		i++
	}
}

func mux(r net.Conn, li string) {
	m = strings.Fields(li)[0] // method
	u = strings.Fields(li)[1] // URL
	fmt.Printf("***METHOD %s\n", m)
	fmt.Printf("***URL %s\n", u)
	// multiplaxer
	if m == "GET" && u == "/" {
		index(r)
	}
	if m == "GET" && u == "/about" {
		about(r)
	}
	if m == "GET" && u == "/contact" {
		contact(r)
	}
	if m == "GET" && u == "/apply" {
		apply(r)
	}
	if m == "POST" && u == "/apply" {
		applyPOST(r)
	}
}

// func response(w net.Conn) {
// 	if m == "GET" && u == "/" {
// 		index(w)
// 	}
// 	if m == "GET" && u == "/about" {
// 		about(w)
// 	}
// 	if m == "GET" && u == "/contact" {
// 		contact(w)
// 	}
// 	if m == "GET" && u == "/apply" {
// 		apply(w)
// 	}
// 	if m == "POST" && u == "/apply" {
// 		applyPOST(w)
// 	}
// }

func index(w net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(w, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(w, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(w, "Content-Type: text/html\r\n")
	fmt.Fprint(w, "\r\n")
	fmt.Fprint(w, body)
}

func about(w net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(w, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(w, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(w, "Content-Type: text/html\r\n")
	fmt.Fprint(w, "\r\n")
	fmt.Fprint(w, body)
}

func contact(w net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(w, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(w, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(w, "Content-Type: text/html\r\n")
	fmt.Fprint(w, "\r\n")
	fmt.Fprint(w, body)
}

func apply(w net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="POST!">
	</form>
	</body></html>`
	fmt.Fprint(w, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(w, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(w, "Content-Type: text/html\r\n")
	fmt.Fprint(w, "\r\n")
	fmt.Fprint(w, body)
}

func applyPOST(w net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(w, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(w, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(w, "Content-Type: text/html\r\n")
	fmt.Fprint(w, "\r\n")
	fmt.Fprint(w, body)
}
