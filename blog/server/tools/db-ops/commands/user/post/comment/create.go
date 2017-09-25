package comment

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports
	"fmt"
	"github.com/hofstadter-io/data-utils/io"
	"github.com/hofstadter-io/examples/blog/lib/types"
	"github.com/hofstadter-io/examples/blog/server/databases/postgres"

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   serverToolDB
// Name:   create
// Usage:  create <post-uuid> <user-uuid> <data-file>
// Parent: comment

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var CreateCmd = &cobra.Command{

	Use: "create <post-uuid> <user-uuid> <data-file>",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In createCmd", "args", args)
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

		// [2]name:   data-file
		//     help:
		//     req'd:  true
		if 2 >= len(args) {
			fmt.Println("missing required argument: 'data-file'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var dataFile string

		if 2 < len(args) {

			dataFile = args[2]
		}

		var comment types.Comment
		var iface interface{}
		iface = &comment
		// unmarshal data file into struct
		_, cerr := io.ReadFile(dataFile, &iface)
		if cerr != nil {
			fmt.Println("Error", cerr)
			os.Exit(1)
		}

		fmt.Printf("%+v\n\n", comment)

		// validate the input here

		err := types.CreateComment(
			postgres.POSTGRES,
			&comment,
			postUUID,
			userUUID,
		)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		fmt.Println("Created:", comment.UUID)

	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
