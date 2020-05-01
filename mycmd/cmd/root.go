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
	app []appcommands
}

type appcommands struct {
	commands []string
}
