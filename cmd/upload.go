package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload [file]",
	Short: "Uploads a package via yaml definition to crackle.pm or configured Crackle endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo")
	},
}

func init() {
	Root.AddCommand(uploadCmd)
}
