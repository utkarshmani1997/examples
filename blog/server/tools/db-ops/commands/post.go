package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/examples/blog/server/tools/db-ops/commands/post"
)

// Tool:   serverToolDB
// Name:   post
// Usage:
// Parent: serverToolDB

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var PostCmd = &cobra.Command{

	Use: "post",

	Short: "work with the post resource",
}

func init() {
	RootCmd.AddCommand(PostCmd)
}

func init() {
	// add sub-commands to this command when present

	PostCmd.AddCommand(post.MigrateCmd)
	PostCmd.AddCommand(post.CreateCmd)
	PostCmd.AddCommand(post.FindCmd)
	PostCmd.AddCommand(post.UpdateCmd)
	PostCmd.AddCommand(post.DeleteCmd)
	PostCmd.AddCommand(post.QueryCmd)
}

// HOFSTADTER_BELOW
