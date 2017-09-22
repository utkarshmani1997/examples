package user

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
// Name:   migrate
// Usage:  migrate
// Parent: user

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var MigrateCmd = &cobra.Command{

	Use: "migrate",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In migrateCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		fmt.Println("serverToolDB user migrate:")
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
