package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add new data,command",
	Example: RootCmd.Use + " add -l github git init",
	Run: func(cmd *cobra.Command, args []string) {
		//*flags --l --c functions
		lAdd(cmd, args)
		cAdd(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().String("l", " ", "adds command to the given list")
	addCmd.Flags().String("c", " ", "creates a list")
}
