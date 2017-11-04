package cmd

import (
	"context"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sunshinekitty/cr/pkg/crackle"
)

const (
	version = "0.0.0"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version of client and configured endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Client: %s\n", version)
		client := crackle.NewClient(nil)
		client.BaseURL, _ = url.Parse(viper.GetString("crackle.api"))
		ctx := context.Background()
		server, resp, _ := client.Version.Server(ctx)
		switch resp.StatusCode {
		case 200:
			fmt.Printf("Server: %s\n", server.Version)
		default:
			fmt.Printf("%v: %s\n", resp.StatusCode, resp.Body)
		}
	},
}

func init() {
	Root.AddCommand(versionCmd)
}
