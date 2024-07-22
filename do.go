package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Println("Marking task as completed:", task)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
