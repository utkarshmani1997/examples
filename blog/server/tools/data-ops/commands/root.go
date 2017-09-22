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

		Use: "server-tool-data",

		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug("In server-tool-dataCmd", "args", args)
			// Argument Parsing

			// HOFSTADTER_START cmd_run
			fmt.Println("server-tool-data:")
			// HOFSTADTER_END   cmd_run
		},
	}
)

/*
ctx_path: dsl.server.tools.data-ops.cli
name: server-tool-data
parent: ""
parent_path: ""
pkg_path: server/tools/data-ops
pkgPath: server-tool-data
relPath: server/tools/data-ops

*/

// HOFSTADTER_BELOW
