package cmd

import (
	"context"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/sunshinekitty/cr/crackle"
	"github.com/sunshinekitty/cr/helpers"
)

var update bool

var getCmd = &cobra.Command{
	Use:   "get [package]",
	Short: "Install tool from Crackle",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Provide package to install")
		}

		if len(args) > 1 || !helpers.ValidPackageName(args[0]) {
			fmt.Println("Invalid package")
		}

		// TODO: manage updates
		//if update {
		//	fmt.Println("gonna update your shit")
		//} else {
		//	fmt.Println("not gonna update your shit")
		//}

		client := crackle.NewClient(nil)
		client.BaseURL, _ = url.Parse("http://localhost:3813/api/")
		ctx := context.Background()
		fmt.Println(client.Package.GetPackage(ctx, args[0]))
	},
}

func init() {
	getCmd.Flags().BoolVarP(&update, "update", "u", false, "Update package to latest available")
	Root.AddCommand(getCmd)
}
