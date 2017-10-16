package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sunshinekitty/cr/helpers"
)

var filePath string

var uploadCmd = &cobra.Command{
	Use:   "upload [file]",
	Short: "Uploads a package via yaml definition to crackle.pm or configured Crackle endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || len(args) > 1 {
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		}
		p, err := helpers.ConfigToPackageToml(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%v\n", p)
	},
}

func init() {
	Root.AddCommand(uploadCmd)
}
