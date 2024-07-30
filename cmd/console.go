package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Console",
	Long:  "============",
	Run:   runConsole,
}

func runConsole(cmd *cobra.Command, args []string) {
	fmt.Println("Running console scripts")
}
