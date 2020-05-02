package cmd

import "github.com/spf13/cobra"

//Data is struct for storing my commands
var Data data

//RootCmd is the main command
var RootCmd = &cobra.Command{
	Use:   "mycmd",
	Short: "mycmd command is for storing my commands",
}

type data struct {
	program []list
}

type list struct {
	name     string
	commands map[string][]string
}
