package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"projects/gophersize/mycmd/cmd"

	"gopkg.in/yaml.v2"
)

//checking if files exist and preaparing file
func init() {

	var (
		err      error
		yamldata []byte
	)

	if cmd.File, err = os.Open(cmd.Path); err != nil {
		cmd.Errorlog.Fatal(err)
	}

	if yamldata, err = ioutil.ReadAll(cmd.File); err != nil {
		cmd.Errorlog.Fatal(err)
	}

	if err = yaml.Unmarshal([]byte(yamldata), cmd.Data); err != nil {
		cmd.Errorlog.Fatal(err)
	}
	fmt.Println(cmd.Data)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
