package cmd

import (

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "TodoApp",
	Aliases: []string{"todo", "task"},
	Short: "A CLI based todo application",
}

func Execute() error {
	return RootCmd.Execute()
}


func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


