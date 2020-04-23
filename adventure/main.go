package main

import (
	"log"
	"net/http"
	"os"
)

type adventure struct {
	info *log.Logger
	err  *log.Logger
	data map[string]topic
}

func main() {
	mylog := &adventure{
		info: log.New(os.Stdout, "Info:\t", log.Ldate|log.Ltime),
		err:  log.New(os.Stdout, "Error:\t", log.Ltime),
		data: parsejson(),
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: mylog,
	}

	mylog.info.Printf("server is listening on port %v", server.Addr)
	server.ListenAndServe()
}
