package main

import "net/http"

func main() {
	//	path := getpath()
	//handler := myhandler(path)
	yamlpath := getyaml()
	yamlhandler := myyamlhandler(yamlpath)
	http.ListenAndServe(":8080", yamlhandler)
}
