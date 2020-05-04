package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

//Yaml struct
type Yaml struct {
	Commands []string `yaml:"commands"`
}

var (
	//Data is val type of struct for storing my commands
	Data = make(map[string]Yaml)
	//File where is stored yaml ,open whenever mycmd command runs
	//TODO implement function for refreshing part
	File *os.File
	//RootCmd is the main command
	RootCmd = &cobra.Command{
		Use:   "mycmd",
		Short: "mycmd command is for storing my commands",
	}
	//Path where to store Yaml
	Path = "data/data.yaml"
	//Infolog to log information
	Infolog = log.New(os.Stdout, "INFO:\t", log.Ltime)
	//Errorlog to log errors
	Errorlog = log.New(os.Stdout, "Error:\t", log.Ltime)
)

func refresh(next func(*cobra.Command, []string, string) error, cmd *cobra.Command, args []string, val string) {
	if err := next(cmd, args, val); err != nil {
		Errorlog.Fatal(err)
	}
	yamldata, err := yaml.Marshal(Data)
	if err != nil {
		Errorlog.Fatal(err)
	}
	if err = ioutil.WriteFile(Path, yamldata, 0666); err != nil {
		Errorlog.Fatal(err)
	}
}

/*

//*Yaml (acutally program_name)
//*command
//*command

cobra
commands:
-cobra add
-cobra help

postgres
commands:
-psql postgres
-\c connect
-\q quit
*/
