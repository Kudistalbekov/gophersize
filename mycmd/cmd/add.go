package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	//Infolog to log information
	Infolog = log.New(os.Stdout, "INFO:", log.Ltime)
	//Errorlog to log errors
	Errorlog = log.New(os.Stdout, "Error:", log.Ltime)
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add new data,command",
	Example: RootCmd.Use + " add -l github git init",
	Run: func(cmd *cobra.Command, args []string) {
		//*flags --l --c functions
		cflagval, _ := cmd.Flags().GetString("c")
		lflagval, _ := cmd.Flags().GetString("l")
		if len(cflagval) != 0 {
			refresh(cAdd, cmd, args, cflagval)
		} else if len(lflagval) != 0 {
			refresh(lAdd, cmd, args, lflagval)
		} else {
			Errorlog.Fatal("flag was not given (--help) to get information")
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().String("l", "", "adds command to the given list")
	addCmd.Flags().String("c", "", "creates a list")
}
