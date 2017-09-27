package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to crackle.pm or configured Crackle endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo")
	},
}

func init() {
	Root.AddCommand(loginCmd)
}
