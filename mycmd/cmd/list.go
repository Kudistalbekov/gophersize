package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show existing list",
	Example: "`list --l git`  or just  `list` ",
	Run: func(cmd *cobra.Command, args []string) {
		flagval, _ := cmd.Flags().GetString("l")
		llist(cmd, args, flagval)

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().String("l", "", "list all or given value's command")
}
