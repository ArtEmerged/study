package main

import (
	"os"
)

func main() {
	name := "Todd McLeod"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	file, _ := os.Create("index.html")

	defer file.Close()
	// io.Copy(file, strings.NewReader(tpl))
	file.WriteString(tpl)
}
