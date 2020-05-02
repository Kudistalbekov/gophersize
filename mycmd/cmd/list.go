package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show existing list",
	Run: func(cmd *cobra.Command, args []string) {
		uservalue := strings.Join(args, " ")
		fmt.Printf("list called:%v", uservalue)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
