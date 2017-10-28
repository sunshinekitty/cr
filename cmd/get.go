package cmd

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sunshinekitty/cr/helpers"
	"github.com/sunshinekitty/cr/pkg/crackle"
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
		client.BaseURL, _ = url.Parse(viper.GetString("crackle.api"))
		ctx := context.Background()
		pkg, resp, err := client.Package.GetPackage(ctx, args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		pkgToml, err := helpers.PackageToPackageToml(pkg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if resp.StatusCode == 200 {
			helpers.EnsureConfigDirs()
			crPackageFile, err := helpers.CreatePackageFiles(pkgToml.Package)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			enc := toml.NewEncoder(crPackageFile)
			err = enc.Encode(pkgToml)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			_ = crPackageFile.Close()
			fmt.Printf("Downloaded config for %s\n", pkgToml.Package)
		} else {
			fmt.Printf("%v: %s\n", resp.StatusCode, resp.Body)
		}
	},
}

func init() {
	getCmd.Flags().BoolVarP(&update, "update", "u", false, "Update package to latest available")
	Root.AddCommand(getCmd)
}
