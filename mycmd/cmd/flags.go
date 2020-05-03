package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

func lAdd(cmd *cobra.Command, args []string, flagval string) (map[string]Yaml, error) {
	usercommand := strings.Join(args, " ")
	if list, ok := Data[flagval]; ok == true {
		Data[flagval] = Yaml{Commands: append(list.Commands, usercommand)}
		Infolog.Println("added!")
		return Data, nil
	}
	return nil, errors.New(flagval + " does not exist")
}

func cAdd(cmd *cobra.Command, args []string, flagval string) (map[string]Yaml, error) {
	if _, ok := Data[flagval]; ok == true {
		return nil, errors.New(flagval + " is already exists")
	}
	Data[flagval] = Yaml{}
	Infolog.Println("created!")
	return Data, nil
}
