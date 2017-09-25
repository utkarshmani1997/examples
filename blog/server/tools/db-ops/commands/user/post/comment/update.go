package comment

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   serverToolDB
// Name:   update
// Usage:  update <post-uuid> <user-uuid> <comment-uuid> <data-file>
// Parent: comment

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var UpdateCmd = &cobra.Command{

	Use: "update <post-uuid> <user-uuid> <comment-uuid> <data-file>",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In updateCmd", "args", args)
		// Argument Parsing
		// [0]name:   post-uuid
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'post-uuid'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var postUUID string

		if 0 < len(args) {

			postUUID = args[0]
		}

		// [1]name:   user-uuid
		//     help:
		//     req'd:  true
		if 1 >= len(args) {
			fmt.Println("missing required argument: 'user-uuid'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var userUUID string

		if 1 < len(args) {

			userUUID = args[1]
		}

		// [2]name:   comment-uuid
		//     help:
		//     req'd:  true
		if 2 >= len(args) {
			fmt.Println("missing required argument: 'comment-uuid'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var commentUUID string

		if 2 < len(args) {

			commentUUID = args[2]
		}

		// [3]name:   data-file
		//     help:
		//     req'd:  true
		if 3 >= len(args) {
			fmt.Println("missing required argument: 'data-file'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var dataFile string

		if 3 < len(args) {

			dataFile = args[3]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("serverToolDB user post comment update:",
			postUUID,

			userUUID,

			commentUUID,

			dataFile,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
