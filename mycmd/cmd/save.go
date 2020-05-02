package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// save.goCmd represents the save.go command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save changes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("save called")
	},
}

func init() {
	RootCmd.AddCommand(saveCmd)
}
