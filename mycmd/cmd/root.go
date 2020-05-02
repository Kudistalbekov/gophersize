package cmd

import "github.com/spf13/cobra"

//RootCmd is the main command
var RootCmd = &cobra.Command{
	Use:   "mycmd",
	Short: "mycmd command is for storing my commands",
}

//Data is struct for storing my commands
var Data = make(map[string]data)

type data struct {
	name     string
	commands []string
}

func refresh(next func(*cobra.Command, []string, string) (map[string]data, error), cmd *cobra.Command, args []string, val string) {
	newdata, err := next(cmd, args, val)
	//TODO Implement to refresh data_yaml evretime when there is changes
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
