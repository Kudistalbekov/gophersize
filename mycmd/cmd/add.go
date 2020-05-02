package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new command",
	Run: func(cmd *cobra.Command, args []string) {
		uservalue := strings.Join(args, " ")
		fmt.Printf("command is added: %s\n", uservalue)
		tolist, err := cmd.Flags().GetString("l")
		if err != nil {
			log.Fatal(err)
		} else if tolist == " " {
			log.Println("--l is empty")
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().String("l", " ", "adds the command to given list")
}
