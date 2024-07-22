package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Println("Adding task:", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
