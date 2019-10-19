package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "gotodo",
	Short: "gotodo is a simple TODO application",
	Long:  `Manage your tasks and be more productive!`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
