package cmd

import (
	"fmt"
	"github.com/pojntfx/go-todo/db"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read all todos",
	Long:  `Read all the todos you have`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(db.Read())
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
