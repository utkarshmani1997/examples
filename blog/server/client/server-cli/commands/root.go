package commands

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports
	// infered imports

	"github.com/spf13/cobra"
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var (
	RootCmd = &cobra.Command{

		Use: "server-cli",

		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug("In server-cliCmd", "args", args)
			// Argument Parsing

			// HOFSTADTER_START cmd_run
			fmt.Println("server-cli:")
			// HOFSTADTER_END   cmd_run
		},
	}
)

/*
ctx_path: dsl.server.client.server-cli.cli
name: server-cli
parent: ""
parent_path: ""
pkg_path: server/client/server-cli
pkgPath: server-cli
relPath: server/client/server-cli

*/

// HOFSTADTER_BELOW
