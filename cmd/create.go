package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/pojntfx/gotodo/db"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a todo",
	Long:  `Create a new todo and track your time!`,
	Run: func(cmd *cobra.Command, args []string) {
		db.ReadFromFile()
		ID := len(db.Read()) + 1
		db.Create(db.Todo{Id: ID, Title: title, Description: description})
		db.WriteToFile()
		fmt.Println("Successfully created gotodo:")
		t := table.NewWriter()
		t.AppendHeader(table.Row{"#", "Title", "Description"})
		t.AppendRow(table.Row{
			ID,
			title,
			description,
		})
		fmt.Println(t.Render())
	},
}

var title string
var description string

func init() {
	createCmd.Flags().StringVarP(&title, "title", "t", "My first gotodo", "The title of the new todo")
	createCmd.Flags().StringVarP(&description, "description", "d", "Buy groceries", "The description of the new todo")

	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("description")

	rootCmd.AddCommand(createCmd)
}
