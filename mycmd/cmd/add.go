package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds new data , command",
	Example: RootCmd.Use + " add --c github",
	Run: func(cmd *cobra.Command, args []string) {
		//*flags --l --c functions
		cflagval, _ := cmd.Flags().GetString("c")
		lflagval, _ := cmd.Flags().GetString("l")
		if len(cflagval) != 0 || len(lflagval) != 0 {

			if len(cflagval) != 0 {
				//run cAdd then refresh the data in file
				refresh(cAdd, cmd, args, cflagval)
			}
			if len(lflagval) != 0 {
				//run lAdd then refresh the data in file
				refresh(lAdd, cmd, args, lflagval)
			}
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
