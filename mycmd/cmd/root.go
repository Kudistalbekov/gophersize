package cmd

import "github.com/spf13/cobra"

//Data is struct for storing my commands
var Data map[string]data

//RootCmd is the main command
var RootCmd = &cobra.Command{
	Use:   "mycmd",
	Short: "mycmd command is for storing my commands",
}

type data struct {
	name     string
	commands []string
}

/*

//*data (acutally program_name)
//*command
//*command

cobra
-cobra add
-cobra help

postgres
-psql postgres
-\c connect
-\q quit
*/
