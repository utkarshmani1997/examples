package commands

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   serverToolDB
// Name:   query
// Usage:
// Parent: serverToolDB

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var QueryCmd = &cobra.Command{

	Use: "query",

	Short: "query records",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In queryCmd", "args", args)
		// Argument Parsing
		// [0]name:   query
		//     help:
		//     req'd:

		var query string

		if 0 < len(args) {

			query = args[0]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("serverToolDB query:",
			query,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(QueryCmd)
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
