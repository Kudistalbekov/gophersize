package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove existing command",
	Run: func(cmd *cobra.Command, args []string) {
		var task []int
		for _, uservalue := range args {
			if n, err := strconv.Atoi(uservalue); err != nil {
				log.Printf("couldn't remove %v", uservalue)
			} else {
				task = append(task, n)
			}
		}
		fmt.Println(task)
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
