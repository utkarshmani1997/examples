package user

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/examples/blog/server/tools/db-ops/commands/user/post"
)

// Tool:   serverToolDB
// Name:   post
// Usage:
// Parent: user

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
	// add sub-commands to this command when present

	PostCmd.AddCommand(post.CommentCmd)
	PostCmd.AddCommand(post.MigrateCmd)
	PostCmd.AddCommand(post.CreateCmd)
	PostCmd.AddCommand(post.FindCmd)
	PostCmd.AddCommand(post.UpdateCmd)
	PostCmd.AddCommand(post.DeleteCmd)
}

// HOFSTADTER_BELOW
