package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/pojntfx/gotodo/db"
	"github.com/spf13/cobra"
	"strconv"
)

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a todo",
	Long:  `Update a todo and improve your time tracking!`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db.ReadFromFile()
		ID, _ := strconv.ParseInt(args[0], 0, 64)
		updatedTodo, err := db.Update(db.Todo{Id: int(ID), Title: title, Description: description})
		if err != nil {
			fmt.Println("Could not update this todo:", err)
		} else {
			db.WriteToFile()
			fmt.Println("Successfully updated gotodo:")
			t := table.NewWriter()
			t.AppendHeader(table.Row{"#", "Title", "Description"})
			t.AppendRow(table.Row{
				updatedTodo.Id,
				updatedTodo.Title,
				updatedTodo.Description,
			})
			fmt.Println(t.Render())
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&title, "title", "t", "", "The new title of the todo to update")
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "The new description of the todo to update")

	rootCmd.AddCommand(updateCmd)
}
