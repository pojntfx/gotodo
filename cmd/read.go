package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/pojntfx/gotodo/db"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read all todos",
	Long:  `Read all the todos you have`,
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()
		t.AppendHeader(table.Row{"#", "Title", "Description"})
		db.ReadFromFile()
		for _, todo := range db.Read() {
			t.AppendRow(table.Row{
				todo.Id,
				todo.Title,
				todo.Description,
			})
		}
		fmt.Println(t.Render())
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
