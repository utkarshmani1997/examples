package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports
	"fmt"
	"github.com/hofstadter-io/examples/blog/server/databases/postgres"
	"os"

	// infered imports
	// infered imports

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var ServerToolDBLong = `A tool for working with the server databases`

var (
	RootConfigPFlag string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&RootConfigPFlag, "config", "C", "server/config.yaml", "the config file for the API server")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))

}

var (
	RootCmd = &cobra.Command{

		Use: "serverToolDB",

		Short: "A tool for working with the server databases",

		Long: ServerToolDBLong,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Argument Parsing

			filename := "server/config.yaml"
			f, err := os.Open(filename)
			if err != nil {
				fmt.Println("Error", err)
				os.Exit(1)
			}
			viper.SetConfigType("yaml")
			verr := viper.MergeConfig(f)
			if verr != nil {
				fmt.Println("Error", verr)
				os.Exit(1)
			}
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

		},

		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			logger.Debug("In PersistentPostRun serverToolDBCmd", "args", args)
			// Argument Parsing

			postgres.DisconnectFromPsql()

		},
	}
)

// HOFSTADTER_BELOW
