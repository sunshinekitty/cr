package cmd

import "github.com/spf13/cobra"

// Root is our command object
var Root = &cobra.Command{
	Use:   "cr",
	Short: "Package manager for container based applications",
}
