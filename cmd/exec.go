package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/codeskyblue/go-sh"
	"github.com/spf13/cobra"

	"github.com/sunshinekitty/cr/helpers"
)

var execCmd = &cobra.Command{
	Use:   "exec [package]",
	Short: "Executes package based on package config",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			exit1("Provide package to exec")
		}

		pkg := args[0]
		if len(args) > 1 || !helpers.ValidPackageName(pkg) {
			exit1("Invalid package")
		}

		configFile := fmt.Sprintf("%s/packages/%s.toml", helpers.ConfigDir(), pkg)

		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			fmt.Printf("Config for package %s doesn't exist\n", pkg)
			exit1("Download a package with `cr get [package]`")
		}

		runCmd, runArgs, err := helpers.ConfigFileToCmd(configFile)
		if err != nil {
			exit1(err.Error())
		}

		sh.Command(runCmd, strings.Split(runArgs, " ")).Run()
	},
}

func init() {
	Root.AddCommand(execCmd)
}
