package cmd

import (
	"fmt"
	"github.com/pojntfx/go-todo/db"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a todo",
	Long:  `Create a new todo and track your time!`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Create(db.Todo{Id: len(db.Read()) + 1, Title: title, Description: description})
		fmt.Println("TODO added successfully!")
	},
}

var title string
var description string

func init() {
	createCmd.Flags().StringVarP(&title, "title", "t", "My first TODO", "The title of the new todo")
	createCmd.Flags().StringVarP(&description, "description", "d", "Buy groceries", "The description of the new todo")

	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("description")

	rootCmd.AddCommand(createCmd)
}
