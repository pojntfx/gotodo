package cmd

import (
	"fmt"
	"github.com/pojntfx/gotodo/db"
	"github.com/spf13/cobra"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <ids>",
	Short: "Delete a todo",
	Long:  `Delete a specific todo`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db.ReadFromFile()
		for _, arg := range args {
			id, _ := strconv.ParseInt(arg, 0, 64)
			db.Delete(int(id))
			fmt.Println("Successfully deleted gotodo", arg)
		}
		db.WriteToFile()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
