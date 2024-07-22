package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "todo list manager",
	Long:  "CMD application to manage your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
