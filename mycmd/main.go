package main

//TODO Make it work with commands below :
//* mycmd remove -l listname
//* mtcmd remove -l listname 0 1 (number of commands)
import (
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
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
