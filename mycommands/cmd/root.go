package cmd

import "github.com/spf13/cobra"

//RootCmd command is the root command
var RootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "cmd command is for storing my cli commands",
}
