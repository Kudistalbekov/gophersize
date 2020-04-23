package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func (a adventure) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url := r.URL.Path
		temp, err := template.ParseFiles("template.html")
		if err != nil {
			a.err.Println(err)
		}
		if v, ok := a.data[strings.Replace(url, "/", "", -1)]; ok {
			temp.Execute(w, v)
		} else {
			http.Error(w, "path does not exist", http.StatusBadRequest)
		}
		return
	}
	http.Error(w, "error method is not GET request", http.StatusBadRequest)
}

func parsejson() map[string]topic {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	var arc map[string]topic
	err = json.Unmarshal(data, &arc)
	if err != nil {
		log.Fatal(err)
	}
	return arc
}
