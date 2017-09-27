package cmd

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sunshinekitty/cr/db"
	"github.com/sunshinekitty/cr/web/handlers"
)

var webCmd = &cobra.Command{
	Use:   "web [host:port]",
	Short: "Starts Crackle web server on host:port (default 0.0.0.0:3813)",
	Run: func(cmd *cobra.Command, args []string) {
		bind := "0.0.0.0:3813"
		if len(args) > 0 {
			bind = args[0]
		}
		// Echo instance
		e := echo.New()

		// App config
		e.HideBanner = true
		e.Logger.SetLevel(log.INFO)

		// Viper config
		viper.SetConfigName("server")
		viper.SetConfigType("toml")
		viper.AddConfigPath("/etc/crackle/")
		viper.AddConfigPath("$HOME/.crackle")
		viper.AddConfigPath(".")
		viper.SetDefault("LogLevel", "info")
		viper.SetDefault("database.MaxConnections", 50)
		viper.AutomaticEnv()
		err := viper.ReadInConfig()
		if err != nil { // Handle errors reading the config file
			e.Logger.Fatal(fmt.Errorf("Fatal error config file: %s", err))
		}
		viper.WatchConfig()
		viper.OnConfigChange(func(fse fsnotify.Event) {
			e.Logger.Info("Config file reloaded: ", fse.Name)
		})
		db.InitDB()

		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		// Route => handler
		e.POST("/api/package/:name", handlers.CreatePackage)
		e.GET("/api/package/:name", handlers.ReadPackage)
		e.PUT("/api/package/:name", handlers.UpdatePackage)
		e.DELETE("/api/package/:name", handlers.DeletePackage)

		// Start server
		e.Logger.Info("Starting server at ", bind)
		e.Logger.Fatal(e.Start(bind))
	},
}

func init() {
	Root.AddCommand(webCmd)
}
