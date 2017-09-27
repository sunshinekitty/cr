package db

import (
	"fmt"
	"log"

	"github.com/sunshinekitty/cr/web/handlers"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// InitDB initializes all database pointers in the project
func InitDB() {
	if viper.GetString("database.driver") == "psql" {
		var err error
		handlers.DB, err = sqlx.Connect("postgres",
			fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
				viper.GetString("database.user"),
				viper.GetString("database.pass"),
				viper.GetString("database.db")))
		if err != nil {
			log.Fatalln(err)
		}

		handlers.DB.SetMaxOpenConns(viper.GetInt("database.MaxConnections"))
	} else {
		log.Fatalln("Invalid database '", viper.GetString("database.driver"), "' specified")
	}
}
