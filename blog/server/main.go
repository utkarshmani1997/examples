package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
	log15 "gopkg.in/inconshreveable/log15.v2"

	"github.com/hofstadter-io/examples/blog/server/resources"

	"github.com/hofstadter-io/examples/blog/server/routes"

	"github.com/hofstadter-io/examples/blog/server/databases/postgres"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// Name:     server
// Version:
// About:    a simple blog server

// HOFSTADTER_START start
// HOFSTADTER_END   start

func main() {
	var err error

	// load the configuration file
	read_config()

	// configure the logger
	config_logger()

	// connect to the database(s)
	postgresHost := viper.GetString("databases.postgres.host")
	postgresPort := viper.GetString("databases.postgres.port")
	postgresUser := viper.GetString("databases.postgres.user")
	postgresPass := viper.GetString("databases.postgres.pass")
	postgresDb := viper.GetString("databases.postgres.dbname")
	postgresSslmode := viper.GetString("databases.postgres.sslmode")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		postgresHost,
		postgresPort,
		postgresUser,
		postgresPass,
		postgresDb,
		postgresSslmode,
	)
	postgres.ConnectToPsql(connStr)
	defer postgres.DisconnectFromPsql()

	// create the echo server object
	E := echo.New()

	// HOFSTADTER_START main-pre-group
	// HOFSTADTER_END   main-pre-group

	// Base API Group
	G := E.Group("api/v1")

	// HOFSTADTER_START main-pre-middleware
	// HOFSTADTER_END   main-pre-middleware

	// Add Middleware
	err = AddMiddleware(G, []string{"default"})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// E.Use(middleware.Recover())
	// HOFSTADTER_START main-pre-routes
	// HOFSTADTER_END   main-pre-routes

	err = resources.InitRouter(G)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = routes.InitRouter(G)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	appHost := viper.GetString("host")
	appPort := viper.GetString("port")

	// HOFSTADTER_START main-prerun
	// HOFSTADTER_END   main-prerun

	E.Logger.SetLevel(log.INFO)
	E.Logger.Fatal(E.Start(appHost + ":" + appPort))
}

func read_config() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.MergeInConfig()

	// Hackery because viper only takes the first config file found... not merging, wtf does merge config mean then anyway
	f, err := os.Open("config.yml")
	if err != nil {
		f = nil
		f2, err2 := os.Open("config.yaml")
		if err2 != nil {
			f = nil
		} else {
			f = f2
		}
	}
	if f != nil {
		verr := viper.MergeConfig(f)
		if verr != nil {
			panic(verr)
		}
	} else {
		panic("missing config.yaml during start up")
	}
}

func config_logger() {
	// log-config default global values
	level := log15.LvlWarn
	stack := false

	// look up in config
	lcfg := viper.GetStringMap("log-config.default")

	if lcfg != nil && len(lcfg) > 0 {
		level_str := lcfg["level"].(string)
		stack = lcfg["stack"].(bool)
		level_local, err := log15.LvlFromString(level_str)
		if err != nil {
			panic(err)
		}
		level = level_local
	}

	termlog := log15.LvlFilterHandler(level, log15.StdoutHandler)
	if stack {
		term_stack := log15.CallerStackHandler("%+v", log15.StdoutHandler)
		termlog = log15.LvlFilterHandler(level, term_stack)
	}

	logger.SetHandler(termlog)

	// set package loggers
	resources.SetLogger(logger)

	routes.SetLogger(logger)

	// HOFSTADTER_START config-logger
	// HOFSTADTER_END   config-logger
}

// HOFSTADTER_BELOW
