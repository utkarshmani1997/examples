package user

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
// Usage:  create <data-file>
// Parent: user

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var CreateCmd = &cobra.Command{

	Use: "create <data-file>",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In createCmd", "args", args)
		// Argument Parsing
		// [0]name:   data-file
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'data-file'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var dataFile string

		if 0 < len(args) {

			dataFile = args[0]
		}

		var user types.User
		var iface interface{}
		iface = &user
		// unmarshal data file into struct
		_, cerr := io.ReadFile(dataFile, &iface)
		if cerr != nil {
			fmt.Println("Error", cerr)
			os.Exit(1)
		}

		fmt.Printf("%+v\n\n", user)

		// validate the input here

		err := types.CreateUser(
			postgres.POSTGRES,
			&user,
		)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		fmt.Println("Created:", user.UUID)

	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
