package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/sunshinekitty/cr/cmd"
)

func main() {
	// Viper config
	viper.SetConfigName("client")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/crackle/")
	viper.AddConfigPath("$HOME/.crackle")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	// We fail silently here since client config isn't needed for web-server
	_ = viper.ReadInConfig()
	if err := cmd.Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
