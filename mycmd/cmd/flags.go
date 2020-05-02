package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

func lAdd(cmd *cobra.Command, args []string, flagval string) (map[string]data, error) {
	usercommand := strings.Join(args, " ")
	if list, ok := Data[flagval]; ok == true {
		Data[flagval] = data{commands: append(list.commands, usercommand)}
		return Data, nil
	}
	return nil, errors.New("list does not exist (create using flag --c)")
}

func cAdd(cmd *cobra.Command, args []string, flagval string) (map[string]data, error) {
	if _, ok := Data[flagval]; ok == true {
		return nil, errors.New("data is already exists")
	}
	Data[flagval] = data{name: flagval}
	return Data, nil
}
