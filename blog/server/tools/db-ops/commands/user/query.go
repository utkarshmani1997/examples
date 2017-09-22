package user

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
// Usage:  query [key=value]...
// Parent: user

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var QueryCmd = &cobra.Command{

	Use: "query [key=value]...",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In queryCmd", "args", args)
		// Argument Parsing
		// [0]name:   matchers
		//     help:
		//     req'd:

		var matchers []string

		if 0 < len(args) {
			matchers = args[0:]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("serverToolDB user query:",
			matchers,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
