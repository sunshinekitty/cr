package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "0.0.0"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version of client and configured endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Client: v%s\n", version)
	},
}

func init() {
	Root.AddCommand(versionCmd)
}
