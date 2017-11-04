package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Root is our command object
var Root = &cobra.Command{
	Use:   "cr",
	Short: "Package manager for container based applications",
}

func exit1(exitString string) {
	fmt.Println(exitString)
	os.Exit(1)
}
