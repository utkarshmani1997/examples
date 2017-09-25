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
// Name:   find
// Usage:  find <post-uuid> <user-uuid> <comment-uuid>
// Parent: comment

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var FindCmd = &cobra.Command{

	Use: "find <post-uuid> <user-uuid> <comment-uuid>",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In findCmd", "args", args)
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

		// HOFSTADTER_START cmd_run
		fmt.Println("serverToolDB user post comment find:",
			postUUID,

			userUUID,

			commentUUID,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
