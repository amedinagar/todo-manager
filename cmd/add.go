package cmd

import (
	"fmt"
	"strings"

	"github.com/amedinagar/todo-manager/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong, unable to include task:", err)
			return
		}
		fmt.Println("Adding \"%s\" task to your list:", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
