package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func lAdd(cmd *cobra.Command, args []string) {
	usercommand := strings.Join(args, " ")
	flagval, err := cmd.Flags().GetString("l")
	if err != nil {
		log.Fatal(err)
	}
	if list, ok := Data[flagval]; ok == true {
		Data[flagval] = data{commands: append(list.commands, usercommand)}
	} else {
		log.Fatal("list does not exist (create using flag --c)")
	}
}

func cAdd(cmd *cobra.Command, args []string) {
	flagval, err := cmd.Flags().GetString("c")
	if err != nil {
		log.Fatal(err)
	}
	if _, ok := Data[flagval]; ok == true {
		log.Fatal("data is already exists")
	} else {
		Data[flagval] = data{}
	}
}
