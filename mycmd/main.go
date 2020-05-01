package main

import (
	"io/ioutil"
	"log"
	"os"
	"projects/gophersize/mycmd/cmd"

	"github.com/go-yaml/yaml"
)

//checking if files exist
func init() {
	file, err := os.Open("data/data.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal([]byte(data), cmd.Data)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
