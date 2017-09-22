package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports
	"github.com/hofstadter-io/examples/blog/lib/types"
	"github.com/hofstadter-io/examples/blog/server/databases/postgres"

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   serverToolDB
// Name:   migrate-tables
// Usage:
// Parent: serverToolDB

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var MigrateTablesCmd = &cobra.Command{

	Use: "migrate-tables",

	Short: "create or migrate all of the tables",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In migrate-tablesCmd", "args", args)
		// Argument Parsing

		types.MigratePostTable(postgres.POSTGRES)

		types.MigrateUserTable(postgres.POSTGRES)

	},
}

func init() {
	RootCmd.AddCommand(MigrateTablesCmd)
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
