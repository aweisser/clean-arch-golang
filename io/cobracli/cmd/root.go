package cmd

import "github.com/spf13/cobra"

var name string

// RootCallback will be called after cli arguments are parsed
var RootCallback func(name string)

// RootCmd provides the cli interface definition
var RootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		RootCallback(name)
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&name, "name", "John Doe", "enter your name")
}
