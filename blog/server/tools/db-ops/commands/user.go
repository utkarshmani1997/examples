package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/examples/blog/server/tools/db-ops/commands/user"
)

// Tool:   serverToolDB
// Name:   user
// Usage:
// Parent: serverToolDB

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var UserCmd = &cobra.Command{

	Use: "user",

	Short: "work with the user resource",
}

func init() {
	RootCmd.AddCommand(UserCmd)
}

func init() {
	// add sub-commands to this command when present

	UserCmd.AddCommand(user.MigrateCmd)
	UserCmd.AddCommand(user.CreateCmd)
	UserCmd.AddCommand(user.FindCmd)
	UserCmd.AddCommand(user.UpdateCmd)
	UserCmd.AddCommand(user.DeleteCmd)
	UserCmd.AddCommand(user.QueryCmd)
}

// HOFSTADTER_BELOW
