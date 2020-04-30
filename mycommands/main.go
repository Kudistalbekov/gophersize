package main

import (
	"log"
	"projects/gophersize/mycommands/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
