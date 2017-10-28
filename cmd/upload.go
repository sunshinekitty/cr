package cmd

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/sunshinekitty/cr/helpers"
	"github.com/sunshinekitty/cr/pkg/crackle"
)

var filePath string

var uploadCmd = &cobra.Command{
	Use:   "upload [file]",
	Short: "Uploads a package via toml definition to crackle.pm or configured Crackle endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || len(args) > 1 {
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		}
		pt, err := helpers.ConfigFileToPackageToml(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err = helpers.ValidPackageToml(pt); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		p, err := helpers.PackageTomlToPackage(pt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client := crackle.NewClient(nil)
		client.BaseURL, _ = url.Parse("http://localhost:3813/api/")
		ctx := context.Background()
		createdPackage, resp, err := client.Package.CreatePackage(ctx, p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if resp.StatusCode == 201 {
			fmt.Printf("Created package %s\n", createdPackage.Name)
		} else {
			fmt.Printf("%v: %s\n", resp.StatusCode, resp.Body)
		}
	},
}

func init() {
	Root.AddCommand(uploadCmd)
}
