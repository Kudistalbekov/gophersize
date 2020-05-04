package cmd

import (
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove existing command",
	Example: "remove --l 3 4 5 6",
	Run: func(cmd *cobra.Command, args []string) {
		flagval, _ := cmd.Flags().GetString("l")
		refresh(lremove, cmd, args, flagval)
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
	removeCmd.Flags().String("l", "", "removes commands from the list")
}
