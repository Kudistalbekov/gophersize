package main

import "fmt"

func main() {
	data := parsehtml("file.html")
	for _, v := range *data {
		fmt.Println(v)
	}
}
